package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/house-hunt-labs/hh-service-property/pkg/models"
)

type ValuationMetricsHandler struct {
	DB *gorm.DB
}

func NewValuationMetricsHandler(db *gorm.DB) *ValuationMetricsHandler {
	return &ValuationMetricsHandler{DB: db}
}

// CreateValuationMetrics creates valuation metrics for a property
func (h *ValuationMetricsHandler) CreateValuationMetrics(c *gin.Context) {
	var metrics models.ValuationMetrics
	if err := c.ShouldBindJSON(&metrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if property exists
	var property models.Property
	if err := h.DB.First(&property, metrics.PropertyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	if err := h.DB.Create(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create valuation metrics"})
		return
	}

	c.JSON(http.StatusCreated, metrics)
}

// GetValuationMetrics retrieves all valuation metrics
func (h *ValuationMetricsHandler) GetValuationMetrics(c *gin.Context) {
	var metrics []models.ValuationMetrics
	if err := h.DB.Find(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch valuation metrics"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

// GetValuationMetricsByProperty retrieves valuation metrics for a specific property
func (h *ValuationMetricsHandler) GetValuationMetricsByProperty(c *gin.Context) {
	propertyID, err := strconv.Atoi(c.Param("propertyId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	var metrics models.ValuationMetrics
	if err := h.DB.First(&metrics, "property_id = ?", propertyID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Valuation metrics not found for this property"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch valuation metrics"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

// UpdateValuationMetrics updates valuation metrics
func (h *ValuationMetricsHandler) UpdateValuationMetrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var metrics models.ValuationMetrics
	if err := h.DB.First(&metrics, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Valuation metrics not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch valuation metrics"})
		return
	}

	if err := c.ShouldBindJSON(&metrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update valuation metrics"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

// DeleteValuationMetrics deletes valuation metrics
func (h *ValuationMetricsHandler) DeleteValuationMetrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.Delete(&models.ValuationMetrics{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete valuation metrics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Valuation metrics deleted"})
}