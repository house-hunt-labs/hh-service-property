# Google Places Data Ingestion Script

This script fetches geographic data from the Google Places API (New) and saves it directly into a MongoDB database.

## Prerequisites

- Python 3.12+
- MongoDB instance
- Google Maps Platform API key with Places API enabled

## Installation

1. **Install Python dependencies:**

```bash
pip install pymongo python-dotenv requests
```

2. **Create a `.env` file** in the `scripts` directory (copy from `.env.example`):

```bash
cp .env.example .env
```

3. **Configure your `.env` file** with the following variables:

```env
# MongoDB Configuration
MONGO_URI=mongodb://localhost:27017
DB_NAME=house_hunt

# Google Maps API Key
Maps_API_KEY=your_google_maps_api_key_here
```

## Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `MONGO_URI` | MongoDB connection string | Yes |
| `DB_NAME` | Database name to use | No (default: "house_hunt") |
| `Maps_API_KEY` | Google Maps Platform API key | Yes |

### Search Queries

Edit the `DEFAULT_SEARCH_QUERIES` list in `ingest_google_places.py` to customize the locations and place types you want to fetch:

```python
DEFAULT_SEARCH_QUERIES = [
    "Hospitals in Whitefield Bangalore",
    "Parks in Whitefield Bangalore",
    # Add more queries...
]
```

Or pass custom queries via command line:

```bash
python ingest_google_places.py "Hospitals in Whitefield" "Parks in Bangalore"
```

## Usage

Run the ingestion script:

```bash
cd scripts
python ingest_google_places.py
```

## What the Script Does

1. **Connects to MongoDB** using the configured URI
2. **Creates indexes** for efficient querying:
   - Unique index on `metadata.google_place_id` for deduplication
   - Unique index on `type` in node_type_styles
3. **Processes each search query:**
   - Fetches places from Google Places API with pagination
   - Normalizes Google types to internal types
   - Ensures corresponding node type styles exist (creates with random color if new)
   - Upserts map nodes (updates existing, inserts new)
4. **Reports statistics** on inserted vs updated documents

## Type Normalization

The script includes a type normalization mapping that converts specific Google types (e.g., `italian_restaurant`, `fast_food_restaurant`) to broader categories (e.g., `restaurant`). You can modify the `TYPE_NORMALIZATION` dictionary in the script to adjust these mappings.

## Database Schema

### node_type_styles Collection

```json
{
  "_id": "ObjectId",
  "type": "string",
  "color": "#RRGGBB",
  "maxRadius": 1000.0,
  "createdAt": "datetime",
  "updatedAt": "datetime"
}
```

### map_nodes Collection

```json
{
  "_id": "ObjectId",
  "type": "string",
  "label": "string",
  "latitude": float64,
  "longitude": float64,
  "description": "string",
  "metadata": {
    "google_place_id": "string",
    "rating": float64,
    "user_ratings_total": int32
  },
  "createdAt": "datetime",
  "updatedAt": "datetime"
}
```

## Notes

- The script uses `upsert` operations based on `google_place_id` to prevent duplicate entries
- There's a 2-second delay between paginated requests to allow the nextPageToken to become valid
- A 1-second delay between different search queries helps avoid rate limiting
- The script creates default node type styles with random hex colors and maxRadius of 1000.0 for any new types encountered