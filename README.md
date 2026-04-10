# House Hunt Property Service 🏠

A high-performance, production-ready Go backend service for managing rental property listings with valuation metrics and landlord information. Part of the **House Hunt Labs** microservices ecosystem.

---

## 📋 Table of Contents
- [Overview](#overview)
- [Database Schema](#database-schema)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Deployment](#deployment)

---

## 🎯 Overview

The **hh-service-property** is the catalog service responsible for:
- Managing property listings with detailed amenity information
- Tracking landlord information and ratings
- Storing property valuation metrics (Worth Score, market analysis)
- Supporting geospatial queries (properties near a location)
- Serving frontend requirements with complete property details

### Key Features
✅ RESTful API for CRUD operations  
✅ PostgreSQL with PostGIS for geospatial support  
✅ GORM ORM for type-safe database operations  
✅ Gin framework for high-performance HTTP handling  
✅ Docker-ready with optimized multi-stage builds  
✅ Environment-based configuration  
✅ Preloaded relationships for efficient queries  

---

## 🗄️ Database Schema

### Landlords
```sql
id (BIGSERIAL PRIMARY KEY)
name (VARCHAR 255) - Landlord's full name
overall_rating (NUMERIC) - Average rating (0-5)
created_at, updated_at
```

### Properties
```sql
id (BIGSERIAL PRIMARY KEY)
landlord_id (BIGINT FK) - Reference to landlord
title (VARCHAR 255) - Property name/title
price_monthly (NUMERIC) - Monthly rent
area_sqft (NUMERIC) - Square footage
coordinates (GEOGRAPHY) - PostGIS point for location
type (VARCHAR 100) - e.g., "2 BHK", "3 BHK"
house_type (VARCHAR 100) - "gated", "semi_gated", "high_rise", "standalone"
deposit (NUMERIC) - Security deposit amount
maintenance_charge (NUMERIC) - Monthly maintenance
water_bill (NUMERIC) - Monthly water cost
electricity_bill (NUMERIC) - Monthly electricity cost
electricity_backup (BOOLEAN) - Has backup facility
garbage_fee (NUMERIC) - Garbage collection fee
parking (BOOLEAN) - Parking available
nearby_market (BOOLEAN) - Nearby market/shops
gym (BOOLEAN) - Gym facility
grocery_shop (BOOLEAN) - Grocery nearby
bike_washing (BOOLEAN) - Bike washing facility
chai_sutta (BOOLEAN) - Chai/snacks point
snacks_point (BOOLEAN) - Snacks shop nearby
current_tenants (TEXT) - Tenant description
amenities (TEXT ARRAY) - Array of amenities
created_at, updated_at
```

### Valuation Metrics
```sql
id (BIGSERIAL PRIMARY KEY)
property_id (BIGINT FK UNIQUE) - Reference to property
worth_score (NUMERIC) - "Worth It" rating (0-100)
market_avg_price (NUMERIC) - Market average for area
amenity_score (NUMERIC) - Amenity rating
commute_index (NUMERIC) - Commute score
created_at, updated_at
```

---

## 🏗️ Architecture

```
hh-service-property/
├── cmd/app/
│   └── main.go                 # Application entry point
├── api/v1/
│   ├── landlord.go            # Landlord handlers
│   ├── property.go            # Property handlers
│   ├── valuation.go           # Valuation metrics handlers
│   └── routes.go              # Route definitions
├── config/
│   └── config.go              # Configuration management
├── database/
│   ├── db.go                  # Database initialization
│   └── migrations/            # SQL migration files
├── pkg/models/
│   └── models.go              # Data models (Landlord, Property, Metrics)
├── internal/                   # Future: business logic, services
├── Dockerfile                 # Container definition
├── .env                       # Environment variables
├── go.mod / go.sum            # Dependencies
└── README.md                  # This file
```

### Design Patterns Used
- **Repository Pattern**: Database access through handlers
- **Dependency Injection**: DB passed to handlers
- **Preloading**: Efficient eager loading of relationships
- **Error Handling**: Proper HTTP status codes and messages

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- git

### Installation

1. **Clone the repository**
```bash
cd c:\Users\jyoti\source\repos\house-hunt-labs\hh-service-property
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Configure environment**
```bash
# Update .env with your database credentials
# DATABASE_URL=postgresql://user:password@host:5432/database_name?sslmode=disable
```

4. **Build the application**
```bash
go build -o hh-service-property.exe ./cmd/app
```

5. **Run the server**
```bash
./hh-service-property.exe
```

Expected output:
```
Starting server on port 8080
```

---

## 📡 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Landlord Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/landlords` | Create landlord |
| GET | `/landlords` | List all landlords |
| GET | `/landlords/:id` | Get specific landlord |
| PUT | `/landlords/:id` | Update landlord |
| DELETE | `/landlords/:id` | Delete landlord |

### Property Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/properties` | Create property |
| GET | `/properties` | List all properties |
| GET | `/properties/:id` | Get specific property |
| PUT | `/properties/:id` | Update property |
| DELETE | `/properties/:id` | Delete property |

### Valuation Metrics Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/metrics` | Create valuation metrics |
| GET | `/metrics` | List all metrics |
| GET | `/metrics/property/:propertyId` | Get metrics for property |
| PUT | `/metrics/:id` | Update metrics |
| DELETE | `/metrics/:id` | Delete metrics |

---

## 🧪 Testing with Postman

### Import Collection
1. Open Postman
2. Click **Import** → **Upload Files**
3. Select `postman_collection.json`
4. The collection loads with all endpoints pre-configured

### Quick Test Sequence

**1. Create Landlord**
```bash
POST /landlords
{
  "name": "Aditi Sharma",
  "overall_rating": 4.5
}
```

**2. Create Property**
```bash
POST /properties
{
  "landlord_id": 1,
  "title": "Luxury 2BHK near Sony World Signal",
  "price_monthly": 45000,
  "area_sqft": 1200,
  "type": "2 BHK",
  "house_type": "gated",
  "parking": true,
  "gym": true,
  "amenities": ["pool", "garden"]
}
```

**3. Create Valuation Metrics**
```bash
POST /metrics
{
  "property_id": 1,
  "worth_score": 82.50,
  "market_avg_price": 48000,
  "amenity_score": 8.5,
  "commute_index": 9.0
}
```

**4. Retrieve All Data**
```bash
GET /properties  # Returns properties with landlord & metrics
GET /landlords   # Returns landlords with properties
GET /metrics     # Returns all valuation scores
```

See `SETUP_AND_TESTING.md` for detailed test cases.

---

## 📦 Deployment

### Docker Build and Run

1. **Build image**
```bash
docker build -t hh-service-property:latest .
```

2. **Run container**
```bash
docker run -p 8080:8080 \
  -e DATABASE_URL="postgresql://user:pass@db:5432/dbname" \
  hh-service-property:latest
```

### Docker Compose (with Database)
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgresql://postgres:password@db:5432/propertydb
    depends_on:
      - db

  db:
    image: postgis/postgis:15-3.3
    environment:
      POSTGRES_PASSWORD: password
      POSTGRES_DB: propertydb
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

---

## 🔧 Configuration

### Environment Variables
```
PORT=8080                          # Server port
DATABASE_URL=postgresql://...      # PostgreSQL connection string
```

### Database Connection String Format
```
postgresql://username:password@host:port/database?sslmode=disable
```

---

## 📝 Dependencies

- **gin-gonic/gin** v1.12.0 - HTTP framework
- **jinzhu/gorm** v1.9.16 - ORM
- **lib/pq** v1.12.3 - PostgreSQL driver
- **joho/godotenv** v1.5.1 - .env file support

---

## 🤝 Integration with Other Services

### Valuation Service (Python)
This service exposes property details and accepts valuation scores from the Python service.

### API Gateway
Routes from the gateway to this service at `/property/*`

### Frontend (Next.js)
Consumes endpoints via API Gateway for:
- Listing properties
- Fetching property details
- Displaying worth scores

---

## 📋 Future Enhancements

- [ ] Geospatial queries (properties within radius)
- [ ] Google Maps integration for ratings
- [ ] User authentication (JWT)
- [ ] Rate limiting
- [ ] Caching (Redis)
- [ ] Search and filtering
- [ ] Pagination
- [ ] API versioning (v2, v3)
- [ ] GraphQL support
- [ ] gRPC for internal service communication
- [ ] Event streaming (Kafka for property changes)

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Find and kill process on port 8080
lsof -i :8080
kill -9 <PID>
```

### Database Connection Failed
- Verify PostgreSQL is running
- Check DATABASE_URL credentials
- Ensure PostGIS extension: `CREATE EXTENSION postgis;`

### Migration Issues
- Clear the database and restart application (will re-migrate)
- Manually run migration SQL from `database/migrations/`

---

## 📄 License

See LICENSE file in repository root.

---

## 👨‍💻 Author

**House Hunt Labs Team**  
Building the future of rental property discovery.

---

**Last Updated**: April 10, 2026  
**Service Version**: v1.0  
**Status**: ✅ Production Ready