package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	propertyHandler := NewPropertyHandler(db)
	landlordHandler := NewLandlordHandler(db)
	metricsHandler := NewValuationMetricsHandler(db)

	v1 := r.Group("/api/v1")
	{
		// Landlord routes
		landlords := v1.Group("/landlords")
		{
			landlords.POST("", landlordHandler.CreateLandlord)
			landlords.GET("", landlordHandler.GetLandlords)
			landlords.GET("/:id", landlordHandler.GetLandlord)
			landlords.PUT("/:id", landlordHandler.UpdateLandlord)
			landlords.DELETE("/:id", landlordHandler.DeleteLandlord)
		}

		// Property routes
		properties := v1.Group("/properties")
		{
			properties.POST("", propertyHandler.CreateProperty)
			properties.GET("", propertyHandler.GetProperties)
			properties.GET("/:id", propertyHandler.GetProperty)
			properties.PUT("/:id", propertyHandler.UpdateProperty)
			properties.DELETE("/:id", propertyHandler.DeleteProperty)
		}

		// Valuation metrics routes
		metrics := v1.Group("/metrics")
		{
			metrics.POST("", metricsHandler.CreateValuationMetrics)
			metrics.GET("", metricsHandler.GetValuationMetrics)
			metrics.GET("/property/:propertyId", metricsHandler.GetValuationMetricsByProperty)
			metrics.PUT("/:id", metricsHandler.UpdateValuationMetrics)
			metrics.DELETE("/:id", metricsHandler.DeleteValuationMetrics)
		}
	}
}