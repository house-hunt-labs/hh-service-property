-- Enable PostGIS extension
CREATE EXTENSION IF NOT EXISTS postgis;

-- Create landlords table
CREATE TABLE IF NOT EXISTS landlords (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    overall_rating NUMERIC DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create properties table with geography support
CREATE TABLE IF NOT EXISTS properties (
    id BIGSERIAL PRIMARY KEY,
    landlord_id BIGINT NOT NULL REFERENCES landlords(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    price_monthly NUMERIC,
    area_sqft NUMERIC,
    coordinates GEOGRAPHY(POINT),
    square_feet NUMERIC,
    _type VARCHAR(100),
    house_type VARCHAR(100),
    deposit NUMERIC,
    maintenance_charge NUMERIC,
    water_bill NUMERIC,
    electricity_bill NUMERIC,
    electricity_backup BOOLEAN DEFAULT FALSE,
    garbage_fee NUMERIC,
    parking BOOLEAN DEFAULT FALSE,
    nearby_market BOOLEAN DEFAULT FALSE,
    gym BOOLEAN DEFAULT FALSE,
    grocery_shop BOOLEAN DEFAULT FALSE,
    bike_washing BOOLEAN DEFAULT FALSE,
    chai_sutta BOOLEAN DEFAULT FALSE,
    snacks_point BOOLEAN DEFAULT FALSE,
    current_tenants TEXT,
    amenities TEXT[] DEFAULT ARRAY[]::TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on coordinates for faster geo queries
CREATE INDEX IF NOT EXISTS idx_properties_coordinates ON properties USING GIST(coordinates);
CREATE INDEX IF NOT EXISTS idx_properties_landlord_id ON properties(landlord_id);

-- Create valuation_metrics table
CREATE TABLE IF NOT EXISTS valuation_metrics (
    id BIGSERIAL PRIMARY KEY,
    property_id BIGINT NOT NULL UNIQUE REFERENCES properties(id) ON DELETE CASCADE,
    worth_score NUMERIC,
    market_avg_price NUMERIC,
    amenity_score NUMERIC,
    commute_index NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_valuation_metrics_property_id ON valuation_metrics(property_id);