package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/house-hunt-labs/hh-service-property/pkg/models"
)

func InitDB(databaseURL string) *gorm.DB {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(
		&models.Landlord{},
		&models.Property{},
		&models.ValuationMetrics{},
	)

	return db
}