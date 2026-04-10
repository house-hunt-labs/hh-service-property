package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-hunt-labs/hh-service-property/config"
	"github.com/house-hunt-labs/hh-service-property/database"
	"github.com/house-hunt-labs/hh-service-property/api/v1"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db := database.InitDB(cfg.DatabaseURL)
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	v1.SetupRoutes(r, db)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}