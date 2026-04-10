package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Landlord represents a property landlord
type Landlord struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"not null"`
	OverallRating  float64        `json:"overall_rating"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Properties     []Property     `json:"properties,omitempty" gorm:"foreignKey:LandlordID"`
}

// Property represents a rental property with geospatial support
type Property struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	LandlordID       uint           `json:"landlord_id" gorm:"not null"`
	Landlord         *Landlord      `json:"landlord,omitempty" gorm:"foreignKey:LandlordID"`
	Title            string         `json:"title" gorm:"not null"`
	PriceMonthly     float64        `json:"price_monthly"`
	AreaSqft         float64        `json:"area_sqft"`
	Coordinates      sql.NullString `json:"coordinates" gorm:"type:geography(point,4326)"`
	SquareFeet       float64        `json:"square_feet"`
	Type             string         `json:"type"` // e.g., "2 BHK"
	HouseType        string         `json:"house_type"` // gated, semi gated, high rise, standalone, single
	Deposit          float64        `json:"deposit"`
	MaintenanceCharge float64       `json:"maintenance_charge"`
	WaterBill        float64        `json:"water_bill"`
	ElectricityBill  float64        `json:"electricity_bill"`
	ElectricityBackup bool          `json:"electricity_backup"`
	GarbageFee       float64        `json:"garbage_fee"`
	Parking          bool           `json:"parking"`
	NearbyMarket     bool           `json:"nearby_market"`
	Gym              bool           `json:"gym"`
	GroceryShop      bool           `json:"grocery_shop"`
	BikeWashing      bool           `json:"bike_washing"`
	ChaiSutta        bool           `json:"chai_sutta"`
	SnacksPoint      bool           `json:"snacks_point"`
	CurrentTenants   string         `json:"current_tenants"`
	Amenities        pq.StringArray `json:"amenities" gorm:"type:text[]"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	ValuationMetrics *ValuationMetrics  `json:"valuation_metrics,omitempty" gorm:"foreignKey:PropertyID"`
}

// ValuationMetrics represents the "Worth It" scores for a property
type ValuationMetrics struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	PropertyID      uint      `json:"property_id" gorm:"not null;uniqueIndex"`
	WorthScore      float64   `json:"worth_score"`
	MarketAvgPrice  float64   `json:"market_avg_price"`
	AmenityScore    float64   `json:"amenity_score"`
	CommuteIndex    float64   `json:"commute_index"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}