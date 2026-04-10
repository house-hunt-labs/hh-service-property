# 🚀 Complete Step-by-Step Setup & Testing Guide

## Overview
This guide walks through setting up the **House Hunt Property Service** and testing it using Postman.

---

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Database Setup](#database-setup)
3. [Environment Configuration](#environment-configuration)
4. [Building the Application](#building-the-application)
5. [Running the Server](#running-the-server)
6. [Testing with Postman](#testing-with-postman)
7. [Complete Test Workflow](#complete-test-workflow)
8. [Troubleshooting](#troubleshooting)

---

## Prerequisites

✅ **Required**:
- Windows 10/11 or Linux/macOS
- Go 1.21 or later
- PostgreSQL 12+ with PostGIS extension
- Postman desktop app or web client
- Git

✅ **Check Installation**:
```powershell
# Check Go version
go version

# Check PostgreSQL version
psql --version
```

---

## Database Setup

### Step 1: Enable PostGIS via Supabase Dashboard

Since you're using **Supabase** (from the `.env` file), PostGIS is already enabled on the `postgres` database.

**Verify PostGIS is enabled:**
```sql
SELECT version();
SELECT PostGIS_version();
```

### Step 2: Create Tables

Connect to your Supabase database and run the migration file:

**Option A: Using Supabase SQL Editor**
1. Go to https://app.supabase.com → Your Project
2. Navigate to **SQL Editor**
3. Click **New Query**
4. Copy the contents from `database/migrations/001_initial_schema.sql`
5. Click **Run**

**Option B: Using psql CLI**
```powershell
psql $env:DATABASE_URL -f database\migrations\001_initial_schema.sql
```

### Step 3: Seed Test Data

Run the seed data to populate test records:

**Option A: Using SQL Editor**
1. Open **SQL Editor** in Supabase
2. Paste contents from `database/migrations/002_seed_data.sql`
3. Click **Run**

**Option B: Using psql CLI**
```powershell
psql $env:DATABASE_URL -f database\migrations\002_seed_data.sql
```

**Expected Result:**
```
INSERT 0 1   (landlord inserted)
INSERT 0 1   (property inserted)
INSERT 0 1   (metrics inserted)
```

---

## Environment Configuration

### Update `.env` File

The `.env` file contains your database connection string from Supabase:

```env
PORT=8080
DATABASE_URL=postgresql://postgres:y8Lg4Lnk$/X?hH-@db.cpljrclbtfendfrffiii.supabase.co:5432/postgres
```

**⚠️ Note**: Keep your DATABASE_URL private! Never commit it to version control.

---

## Building the Application

### Step 1: Navigate to Project Directory
```powershell
cd c:\Users\jyoti\source\repos\house-hunt-labs\hh-service-property
```

### Step 2: Download Dependencies
```powershell
go mod tidy
```

Expected output:
```
go: downloading github.com/gin-gonic/gin v1.12.0
go: downloading github.com/jinzhu/gorm v1.9.16
...
```

### Step 3: Build the Executable
```powershell
go build -o hh-service-property.exe ./cmd/app
```

Expected output:
```
(no output = success)
```

Verify build artifact:
```powershell
Get-ChildItem hh-service-property.exe
```

---

## Running the Server

### Step 1: Start the Application
```powershell
.\hh-service-property.exe
```

**Expected Output:**
```
Starting server on port 8080
```

### Step 2: Verify Server is Running
Open your browser or terminal and fetch:
```powershell
Invoke-WebRequest http://localhost:8080/api/v1/properties
```

Expected response:
```json
[]
```

**✅ Server is running successfully!**

---

## Testing with Postman

### Step 1: Install & Open Postman
- Download: https://www.postman.com/downloads/
- Open the application

### Step 2: Import Collection

**Option A: Import JSON File**
1. Click **Import** button (top left)
2. Select **Upload Files**
3. Choose `postman_collection.json` from project root
4. Click **Import**

**Option B: Create Requests Manually** (described below)

### Step 3: Set Base URL Variable
1. Click **Environments** (left sidebar)
2. Click **+ Create new**
3. Set name: `House Hunt Dev`
4. Add variable:
   - Key: `base_url`
   - Value: `http://localhost:8080/api/v1`
5. Click **Save**

---

## Complete Test Workflow

### Test 1: Create a Landlord ✅

**Request Details:**
- **Method**: POST
- **URL**: `{{base_url}}/landlords`
- **Headers**: `Content-Type: application/json`

**Request Body:**
```json
{
  "name": "Aditi Sharma",
  "overall_rating": 4.5
}
```

**In Postman:**
1. Create new request
2. Set method to **POST**
3. Enter URL: `http://localhost:8080/api/v1/landlords`
4. Go to **Body** tab → select **raw** → **JSON**
5. Paste the request body
6. Click **Send**

**Expected Response (201 Created):**
```json
{
  "id": 1,
  "name": "Aditi Sharma",
  "overall_rating": 4.5,
  "properties": null,
  "created_at": "2026-04-10T10:30:45Z",
  "updated_at": "2026-04-10T10:30:45Z"
}
```

✅ **Note** the `id: 1` for the next tests

---

### Test 2: Get All Landlords ✅

**Request Details:**
- **Method**: GET
- **URL**: `{{base_url}}/landlords`

**In Postman:**
1. Create new request
2. Set method to **GET**
3. Enter URL: `http://localhost:8080/api/v1/landlords`
4. Click **Send**

**Expected Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Aditi Sharma",
    "overall_rating": 4.5,
    "properties": null,
    "created_at": "2026-04-10T10:30:45Z",
    "updated_at": "2026-04-10T10:30:45Z"
  }
]
```

---

### Test 3: Create a Property ✅

**Request Details:**
- **Method**: POST
- **URL**: `{{base_url}}/properties`
- **Headers**: `Content-Type: application/json`

**Request Body:**
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
  "amenities": ["pool", "garden", "play_area"]
}
```

**In Postman:**
1. Create new POST request
2. URL: `http://localhost:8080/api/v1/properties`
3. Body → raw → JSON
4. Paste request body
5. Click **Send**

**Expected Response (201 Created):**
```json
{
  "id": 1,
  "landlord_id": 1,
  "landlord": null,
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 45000,
  "area_sqft": 1200,
  "type": "2 BHK",
  "house_type": "gated",
  "parking": true,
  "gym": true,
  "amenities": ["pool", "garden", "play_area"],
  "created_at": "2026-04-10T10:35:20Z",
  "updated_at": "2026-04-10T10:35:20Z"
}
```

✅ **Note** the `id: 1` for the next tests

---

### Test 4: Create Valuation Metrics ✅

**Request Details:**
- **Method**: POST
- **URL**: `{{base_url}}/metrics`
- **Headers**: `Content-Type: application/json`

**Request Body:**
```json
{
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000.00,
  "amenity_score": 8.5,
  "commute_index": 9.0
}
```

**In Postman:**
1. Create new POST request
2. URL: `http://localhost:8080/api/v1/metrics`
3. Body → raw → JSON
4. Paste request body
5. Click **Send**

**Expected Response (201 Created):**
```json
{
  "id": 1,
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000,
  "amenity_score": 8.5,
  "commute_index": 9,
  "created_at": "2026-04-10T10:40:15Z",
  "updated_at": "2026-04-10T10:40:15Z"
}
```

---

### Test 5: Get All Properties with Relations ✅

**Request Details:**
- **Method**: GET
- **URL**: `{{base_url}}/properties`

**In Postman:**
1. Create new GET request
2. URL: `http://localhost:8080/api/v1/properties`
3. Click **Send**

**Expected Response (200 OK):**
```json
[
  {
    "id": 1,
    "landlord_id": 1,
    "landlord": {
      "id": 1,
      "name": "Aditi Sharma",
      "overall_rating": 4.5,
      "properties": null,
      "created_at": "2026-04-10T10:30:45Z",
      "updated_at": "2026-04-10T10:30:45Z"
    },
    "title": "Luxury 2BHK near Sony World Signal",
    "price_monthly": 45000,
    "area_sqft": 1200,
    "type": "2 BHK",
    "house_type": "gated",
    "parking": true,
    "gym": true,
    "amenities": ["pool", "garden", "play_area"],
    "valuation_metrics": {
      "id": 1,
      "property_id": 1,
      "worth_score": 82.50,
      "market_avg_price": 48000,
      "amenity_score": 8.5,
      "commute_index": 9,
      "created_at": "2026-04-10T10:40:15Z",
      "updated_at": "2026-04-10T10:40:15Z"
    },
    "created_at": "2026-04-10T10:35:20Z",
    "updated_at": "2026-04-10T10:35:20Z"
  }
]
```

---

### Test 6: Get Single Property ✅

**Request Details:**
- **Method**: GET
- **URL**: `{{base_url}}/properties/1`

**In Postman:**
1. Create new GET request
2. URL: `http://localhost:8080/api/v1/properties/1`
3. Click **Send**

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "landlord_id": 1,
  "landlord": { ... },
  "title": "Luxury 2BHK near Sony World Signal",
  ...
}
```

---

### Test 7: Update Property ✅

**Request Details:**
- **Method**: PUT
- **URL**: `{{base_url}}/properties/1`

**Request Body:**
```json
{
  "price_monthly": 47000,
  "maintenance_charge": 600
}
```

**In Postman:**
1. Create new PUT request
2. URL: `http://localhost:8080/api/v1/properties/1`
3. Body → raw → JSON
4. Paste request body
5. Click **Send**

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "price_monthly": 47000,
  "maintenance_charge": 600,
  "updated_at": "2026-04-10T10:45:30Z"
}
```

---

### Test 8: Get Metrics for Property ✅

**Request Details:**
- **Method**: GET
- **URL**: `{{base_url}}/metrics/property/1`

**In Postman:**
1. Create new GET request
2. URL: `http://localhost:8080/api/v1/metrics/property/1`
3. Click **Send**

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000,
  "amenity_score": 8.5,
  "commute_index": 9,
  "created_at": "2026-04-10T10:40:15Z",
  "updated_at": "2026-04-10T10:40:15Z"
}
```

---

### Test 9: Delete Property ✅

**Request Details:**
- **Method**: DELETE
- **URL**: `{{base_url}}/properties/1`

**In Postman:**
1. Create new DELETE request
2. URL: `http://localhost:8080/api/v1/properties/1`
3. Click **Send**

**Expected Response (200 OK):**
```json
{
  "message": "Property deleted"
}
```

---

## Troubleshooting

### Issue 1: Port 8080 Already in Use
```powershell
# Find process using port 8080
Get-NetTCPConnection -LocalPort 8080

# Kill the process
Stop-Process -Id <PID> -Force
```

### Issue 2: Database Connection Error
```
Failed to connect to database: dial tcp db.cpljrclbtfendfrffiii.supabase.co:5432: connection refused
```

**Solution:**
1. Verify DATABASE_URL is correct in `.env`
2. Check your internet connection
3. Verify Supabase project is active
4. Test connection manually with psql

### Issue 3: PostGIS Not Found
```
ERROR: function ST_GeographyFromText does not exist
```

**Solution:**
- Ensure PostGIS is enabled in Supabase
- Run migration files through SQL Editor

### Issue 4: Build Fails
```powershell
# Clean and rebuild
go clean
go mod tidy
go build -o hh-service-property.exe ./cmd/app
```

---

## Success Checklist

✅ PostgreSQL set up with PostGIS  
✅ Database tables created via migrations  
✅ Test data seeded  
✅ `.env` configured with correct DATABASE_URL  
✅ Application builds successfully  
✅ Server runs on port 8080  
✅ All 9 Postman tests pass  
✅ Database records persist correctly  

---

## Next Steps

1. **Set up API Gateway** (Kong/Nginx) to route requests
2. **Deploy Valuation Service** (Python) for Worth Score calculation
3. **Build Frontend** (Next.js) to consume these APIs
4. **Add Authentication** (JWT tokens)
5. **Implement Caching** (Redis)
6. **Add Rate Limiting**
7. **Set up Monitoring** (Prometheus, Grafana)
8. **Configure CI/CD** (GitHub Actions)

---

## Support

For issues or questions:
1. Check `SETUP_AND_TESTING.md` for detailed API documentation
2. Review `README.md` for architecture overview
3. Check error logs in server output
4. Verify database schema: `\dt` in psql

---

**Last Updated**: April 10, 2026  
**Status**: ✅ Ready for Development & Testing