"""
Google Places Data Ingestion Script
Fetches geographic data from Google Places API and saves to MongoDB
"""

import os
import random
import time
from datetime import datetime, timezone
from typing import Optional

import requests
from dotenv import load_dotenv
from pymongo import MongoClient
from pymongo.errors import DuplicateKeyError
from scripts.utils import normalize_type, generate_random_color, get_max_radius

# Load environment variables
load_dotenv()

# Configuration
MONGO_URI = os.getenv("MONGO_URI")
DB_NAME = os.getenv("DB_NAME", "house_hunt")
MAPS_API_KEY = os.getenv("Maps_API_KEY")

# Google Places API configuration
PLACES_API_BASE_URL = "https://places.googleapis.com/v1/places:searchText"

def get_mongo_client() -> MongoClient:
    """Connect to MongoDB using the URI from environment variables."""
    if not MONGO_URI:
        raise ValueError("MONGO_URI not found in environment variables")
    return MongoClient(MONGO_URI)


def get_database(db_client: MongoClient):
    """Get the database instance."""
    return db_client[DB_NAME]


def ensure_node_type_style(db, node_type: str) -> None:
    collection = db["node_type_styles"]

    if collection.find_one({"type": node_type}):
        return

    now = datetime.now(timezone.utc)

    collection.insert_one({
        "type": node_type,
        "color": generate_random_color(),
        "maxRadius": get_max_radius(node_type),
        "createdAt": now,
        "updatedAt": now
    })

    print(f"Created style for {node_type}")



def search_places(query: str, page_token: Optional[str] = None, max_retries: int = 3) -> dict:
    """
    Search for places using Google Places API (New).
    Uses SearchTextRequest with proper field mask.
    Includes retry logic for transient errors.
    """
    headers = {
        "Content-Type": "application/json",
        "X-Goog-Api-Key": MAPS_API_KEY,
        "X-Goog-FieldMask": "places.id,places.displayName,places.location,places.formattedAddress,places.primaryType,places.rating,places.userRatingCount"
    }
    
    payload = {
        "textQuery": query,
        "pageSize": 20
    }
    
    if page_token:
        payload["pageToken"] = page_token
    
    for attempt in range(max_retries):
        try:
            response = requests.post(PLACES_API_BASE_URL, headers=headers, json=payload, timeout=30)
            
            # Print response for debugging
            if response.status_code == 403:
                print(f"  API Error 403: Forbidden - API key may not have Places API enabled or quota exceeded")
                print(f"  Response: {response.text}")
                return {"places": [], "nextPageToken": None}
            elif response.status_code == 429:
                # Rate limited - wait and retry
                wait_time = (attempt + 1) * 5
                print(f"  Rate limited (429), waiting {wait_time}s before retry...")
                time.sleep(wait_time)
                continue
            elif response.status_code != 200:
                print(f"  API Error {response.status_code}: {response.text}")
                response.raise_for_status()

            print(f"  API request successful for query: '{query}' (page token: {page_token})")
            return response.json()
            
        except requests.exceptions.RequestException as e:
            if attempt < max_retries - 1:
                wait_time = (attempt + 1) * 2
                print(f"  Request error: {e}, retrying in {wait_time}s...")
                time.sleep(wait_time)
            else:
                print(f"  Final request error: {e}")
                raise
    
    return {"places": [], "nextPageToken": None}


def extract_place_data(place: dict) -> Optional[dict]:
    """
    Extract and transform place data into our MongoDB schema.
    Returns None if required fields are missing.
    """
    try:
        location = place.get("location", {})
        display_name = place.get("displayName", {})
        
        # Skip places without a valid ID
        google_place_id = place.get("id", "")
        if not google_place_id:
            return None
        
        # Get the primary type and normalize it
        google_type = place.get("primaryType", "")
        normalized_type = normalize_type(google_type)
        
        return {
            "type": normalized_type,
            "label": display_name.get("text", ""),
            "latitude": location.get("latitude", 0.0),
            "longitude": location.get("longitude", 0.0),
            "description": place.get("formattedAddress", ""),
            "metadata": {
                "google_place_id": google_place_id,
                "rating": place.get("rating", 0.0),
                "user_ratings_total": place.get("userRatingCount", 0)
            }
        }
    except Exception as e:
        print(f"Error extracting place data: {e}")
        return None


def upsert_map_node(db, place_data: dict) -> bool:
    """
    Upsert a map node into the database using google_place_id for deduplication.
    Returns True if a new document was inserted, False if updated.
    """
    collection = db["map_nodes"]
    
    google_place_id = place_data["metadata"]["google_place_id"]
    
    # Use upsert to insert or update based on google_place_id
    result = collection.update_one(
        {"metadata.google_place_id": google_place_id},
        {
            "$set": {
                **place_data,
                "updatedAt": datetime.now(timezone.utc)
            },
            "$setOnInsert": {
                "createdAt": datetime.now(timezone.utc)
            }
        },
        upsert=True
    )
    
    return result.upserted_id is not None


def create_indexes(db) -> None:
    """
    Create necessary indexes for the collections.
    Uses partial filter to only index documents where google_place_id is not null.
    """
    # Drop the old index if it exists (without partial filter) to avoid conflicts
    try:
        db["map_nodes"].drop_index("google_place_id_unique")
        print("Dropped old google_place_id_unique index")
    except Exception:
        pass  # Index might not exist
    
    # Unique index on google_place_id in map_nodes (only for non-null values)
    try:
        db["map_nodes"].create_index(
            [("metadata.google_place_id", 1)],
            unique=True,
            name="google_place_id_unique",
            partialFilterExpression={"metadata.google_place_id": {"$exists": True, "$ne": None}}
        )
    except Exception as e:
        # Index might already exist with same definition
        if "already exists" not in str(e).lower():
            print(f"Warning: Could not create google_place_id index: {e}")
    
    # Index on type for faster lookups
    try:
        db["node_type_styles"].create_index(
            [("type", 1)],
            unique=True,
            name="type_unique"
        )
    except Exception as e:
        if "already exists" not in str(e).lower():
            print(f"Warning: Could not create type index: {e}")
    
    print("Indexes created successfully")


def run_ingestion(search_queries: list) -> None:
    """
    Main ingestion function that processes all search queries.
    """
    print("Starting data ingestion...")
    
    # Connect to MongoDB
    client = get_mongo_client()
    db = get_database(client)
    
    try:
        # Create indexes
        # create_indexes(db)
        
        total_inserted = 0
        total_updated = 0
        
        for query in search_queries:
            print(f"\nProcessing query: {query}")
            page_token = None
            page_count = 0
            
            while True:
                try:
                    # Search places
                    results = search_places(query, page_token)
                    places = results.get("places", [])
                    page_token = results.get("nextPageToken")
                    page_count += 1
                    
                    if not places:
                        print(f"  No more results for this query")
                        break
                    
                    print(f"  Found {len(places)} places (page {page_count})")
                    
                    for place in places:
                        place_data = extract_place_data(place)
                        if not place_data:
                            continue
                        
                        # Ensure node type style exists
                        ensure_node_type_style(db, place_data["type"])
                        
                        # Upsert the map node
                        is_new = upsert_map_node(db, place_data)
                        if is_new:
                            total_inserted += 1
                        else:
                            total_updated += 1
                    
                    # If no next page token, break the loop
                    if not page_token:
                        break
                    
                    # Wait before next request to allow token to become valid
                    time.sleep(2)
                    
                except Exception as e:
                    print(f"  Error processing page: {e}")
                    break
            
            # Small delay between queries to avoid rate limiting
            time.sleep(1)
        
        print(f"\n=== Ingestion Complete ===")
        print(f"Total new documents inserted: {total_inserted}")
        print(f"Total documents updated: {total_updated}")
        
    finally:
        client.close()
        print("MongoDB connection closed")

# "Hospitals in Whitefield Bangalore",
# "Parks in Whitefield Bangalore",
# "Schools in Whitefield Bangalore",
# "Shopping malls in Whitefield Bangalore",
# "Restaurants in Whitefield Bangalore",
# "Metro stations in Whitefield Bangalore",
# "Banks in Whitefield Bangalore",
# "Pharmacies in Whitefield Bangalore",
# "Gyms in Whitefield Bangalore",
# "Hotels in Whitefield Bangalore",

# Default search queries - modify as needed
DEFAULT_SEARCH_QUERIES = [
    "Gym in HSR Layout Bangalore",
]


if __name__ == "__main__":
    # You can customize the search queries here
    search_queries = DEFAULT_SEARCH_QUERIES
    
    # Or pass custom queries via command line arguments
    import sys
    if len(sys.argv) > 1:
        search_queries = sys.argv[1:]
    
    run_ingestion(search_queries)