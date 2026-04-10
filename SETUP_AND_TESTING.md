# HouseHunt Property Service - Setup & Testing Guide

## Database Schema Overview

The application uses three main tables:

### 1. **Landlords Table**
```
id (BIGSERIAL PRIMARY KEY)
name (VARCHAR 255, NOT NULL)
overall_rating (NUMERIC)
created_at (TIMESTAMP)
updated_at (TIMESTAMP)
```

### 2. **Properties Table**
```
id (BIGSERIAL PRIMARY KEY)
landlord_id (BIGINT, FOREIGN KEY → landlords.id)
title (VARCHAR 255, NOT NULL)
price_monthly (NUMERIC)
area_sqft (NUMERIC)
coordinates (GEOGRAPHY/PostGIS)
square_feet (NUMERIC)
type (VARCHAR 100) - e.g., "2 BHK"
house_type (VARCHAR 100) - e.g., "gated", "semi_gated", "high_rise"
deposit (NUMERIC)
maintenance_charge (NUMERIC)
water_bill (NUMERIC)
electricity_bill (NUMERIC)
electricity_backup (BOOLEAN)
garbage_fee (NUMERIC)
parking (BOOLEAN)
nearby_market (BOOLEAN)
gym (BOOLEAN)
grocery_shop (BOOLEAN)
bike_washing (BOOLEAN)
chai_sutta (BOOLEAN)
snacks_point (BOOLEAN)
current_tenants (TEXT)
amenities (TEXT ARRAY) - e.g., ["pool", "garden", "play_area"]
created_at (TIMESTAMP)
updated_at (TIMESTAMP)
```

### 3. **Valuation Metrics Table**
```
id (BIGSERIAL PRIMARY KEY)
property_id (BIGINT, UNIQUE FOREIGN KEY → properties.id)
worth_score (NUMERIC) - "Worth It" rating
market_avg_price (NUMERIC)
amenity_score (NUMERIC)
commute_index (NUMERIC)
created_at (TIMESTAMP)
updated_at (TIMESTAMP)
```

---

## Step-by-Step Setup Instructions

### Step 1: Verify PostgreSQL and PostGIS
1. Ensure PostgreSQL is running on your system
2. Verify PostGIS extension is enabled on the database:
   ```sql
   CREATE EXTENSION postgis;
   ```

### Step 2: Update Environment Variables
Edit `.env` file with your database credentials:
```
PORT=8080
DATABASE_URL=postgresql://<username>:<password>@<host>:<port>/<database_name>?sslmode=disable
```

For Supabase (like in your current setup), the URL is already provided.

### Step 3: Build the Application
```bash
cd c:\Users\jyoti\source\repos\house-hunt-labs\hh-service-property
go mod tidy
go build -o hh-service-property.exe ./cmd/app
```

### Step 4: Run the Application
```bash
./hh-service-property.exe
```

You should see:
```
Starting server on port 8080
```

### Step 5: Verify Server is Running
Open your browser and navigate to:
```
http://localhost:8080/api/v1/properties
```

You should get a JSON response (empty array initially).

---

## Postman Testing Guide

### Prerequisites
1. Download &install [Postman](https://www.postman.com/downloads/)
2. Open Postman
3. Create a new collection called "House Hunt Property Service"

### Base URL
```
http://localhost:8080/api/v1
```

---

## Test Cases

### 1. **Create a Landlord**
- **Method**: POST
- **URL**: `http://localhost:8080/api/v1/landlords`
- **Headers**: 
  ```
  Content-Type: application/json
  ```
- **Body** (JSON):
```json
{
  "name": "Aditi Sharma",
  "overall_rating": 4.5
}
```
- **Expected Response** (201 Created):
```json
{
  "id": 1,
  "name": "Aditi Sharma",
  "overall_rating": 4.5,
  "created_at": "2026-04-10T10:30:00Z",
  "updated_at": "2026-04-10T10:30:00Z"
}
```

---

### 2. **Get All Landlords**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/landlords`
- **Expected Response** (200 OK):
```json
[
  {
    "id": 1,
    "name": "Aditi Sharma",
    "overall_rating": 4.5,
    "properties": [],
    "created_at": "2026-04-10T10:30:00Z",
    "updated_at": "2026-04-10T10:30:00Z"
  }
]
```

---

### 3. **Get Single Landlord**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/landlords/1`
- **Expected Response** (200 OK):
```json
{
  "id": 1,
  "name": "Aditi Sharma",
  "overall_rating": 4.5,
  "properties": [],
  "created_at": "2026-04-10T10:30:00Z",
  "updated_at": "2026-04-10T10:30:00Z"
}
```

---

### 4. **Create a Property**
- **Method**: POST
- **URL**: `http://localhost:8080/api/v1/properties`
- **Headers**: 
  ```
  Content-Type: application/json
  ```
- **Body** (JSON):
```json
{
  "landlord_id": 1,
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 45000,
  "area_sqft": 1200,
  "type": "2 BHK",
  "house_type": "gated",
  "square_feet": 1200,
  "deposit": 90000,
  "maintenance_charge": 500,
  "water_bill": 2000,
  "electricity_bill": 3000,
  "electricity_backup": true,
  "garbage_fee": 100,
  "parking": true,
  "nearby_market": true,
  "gym": true,
  "grocery_shop": true,
  "bike_washing": true,
  "chai_sutta": false,
  "snacks_point": true,
  "current_tenants": "Young professionals",
  "amenities": ["pool", "garden", "play_area"],
  "coordinates": {"type": "Point", "coordinates": [77.6245, 12.9339]}
}
```
- **Expected Response** (201 Created):
```json
{
  "id": 1,
  "landlord_id": 1,
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 45000,
  "area_sqft": 1200,
  "type": "2 BHK",
  "house_type": "gated",
  "parking": true,
  "gym": true,
  "created_at": "2026-04-10T10:35:00Z",
  "updated_at": "2026-04-10T10:35:00Z"
}
```

---

### 5. **Get All Properties**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/properties`
- **Expected Response** (200 OK):
```json
[
  {
    "id": 1,
    "landlord_id": 1,
    "landlord": {
      "id": 1,
      "name": "Aditi Sharma",
      "overall_rating": 4.5
    },
    "title": "Luxury 2BHK near Sony World Signal",
    "price_monthly": 45000,
    "area_sqft": 1200,
    "type": "2 BHK",
    "valuation_metrics": null,
    "created_at": "2026-04-10T10:35:00Z"
  }
]
```

---

### 6. **Get Single Property**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/properties/1`
- **Expected Response** (200 OK):
```json
{
  "id": 1,
  "landlord_id": 1,
  "landlord": {
    "id": 1,
    "name": "Aditi Sharma",
    "overall_rating": 4.5
  },
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 45000,
  "area_sqft": 1200,
  "type": "2 BHK",
  "house_type": "gated",
  "parking": true,
  "valuation_metrics": null,
  "created_at": "2026-04-10T10:35:00Z"
}
```

---

### 7. **Create Valuation Metrics for a Property**
- **Method**: POST
- **URL**: `http://localhost:8080/api/v1/metrics`
- **Headers**: 
  ```
  Content-Type: application/json
  ```
- **Body** (JSON):
```json
{
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000.00,
  "amenity_score": 8.5,
  "commute_index": 9.0
}
```
- **Expected Response** (201 Created):
```json
{
  "id": 1,
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000.00,
  "amenity_score": 8.5,
  "commute_index": 9.0,
  "created_at": "2026-04-10T10:40:00Z",
  "updated_at": "2026-04-10T10:40:00Z"
}
```

---

### 8. **Get All Valuation Metrics**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/metrics`
- **Expected Response** (200 OK):
```json
[
  {
    "id": 1,
    "property_id": 1,
    "worth_score": 82.50,
    "market_avg_price": 48000.00,
    "amenity_score": 8.5,
    "commute_index": 9.0,
    "created_at": "2026-04-10T10:40:00Z",
    "updated_at": "2026-04-10T10:40:00Z"
  }
]
```

---

### 9. **Get Metrics for Specific Property**
- **Method**: GET
- **URL**: `http://localhost:8080/api/v1/metrics/property/1`
- **Expected Response** (200 OK):
```json
{
  "id": 1,
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000.00,
  "amenity_score": 8.5,
  "commute_index": 9.0,
  "created_at": "2026-04-10T10:40:00Z",
  "updated_at": "2026-04-10T10:40:00Z"
}
```

---

### 10. **Update a Property**
- **Method**: PUT
- **URL**: `http://localhost:8080/api/v1/properties/1`
- **Headers**: 
  ```
  Content-Type: application/json
  ```
- **Body** (JSON):
```json
{
  "price_monthly": 47000,
  "maintenance_charge": 600,
  "gym": false
}
```
- **Expected Response** (200 OK):
```json
{
  "id": 1,
  "landlord_id": 1,
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 47000,
  "maintenance_charge": 600,
  "gym": false,
  "updated_at": "2026-04-10T10:45:00Z"
}
```

---

### 11. **Delete a Property**
- **Method**: DELETE
- **URL**: `http://localhost:8080/api/v1/properties/1`
- **Expected Response** (200 OK):
```json
{
  "message": "Property deleted"
}
```

---

### 12. **Delete a Landlord**
- **Method**: DELETE
- **URL**: `http://localhost:8080/api/v1/landlords/1`
- **Expected Response** (200 OK):
```json
{
  "message": "Landlord deleted"
}
```

---

## Common Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid ID"
}
```

### 404 Not Found
```json
{
  "error": "Property not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to create property"
}
```

---

## Docker Setup (Optional)

### Build Docker Image
```bash
docker build -t hh-service-property .
```

### Run Container
```bash
docker run -p 8080:8080 \
  -e DATABASE_URL="postgresql://user:password@db:5432/propertydb?sslmode=disable" \
  hh-service-property
```

---

## Troubleshooting

1. **Connection refused on port 8080**
   - Check if previous instance is running: `netstat -ano | findstr :8080`
   - Kill if necessary: `taskkill /PID <PID> /F`

2. **Database connection error**
   - Verify PostgreSQL is running
   - Check DATABASE_URL in .env
   - Test connection: `psql <database_url>`

3 **PostGIS error**
   - Ensure PostGIS extension exists: `CREATE EXTENSION postgis;`
   - Verify in database: `SELECT * FROM pg_extension;`

---

## Next Steps
- Integrate with Valuation Service (Python FastAPI)
- Add authentication (JWT)
- Implement geospatial queries (properties near user)
- Add rate limiting
- Create monitoring/logging