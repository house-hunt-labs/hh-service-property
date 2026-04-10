package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/house-hunt-labs/hh-service-property/pkg/models"
)

type PropertyHandler struct {
	DB *gorm.DB
}

func NewPropertyHandler(db *gorm.DB) *PropertyHandler {
	return &PropertyHandler{DB: db}
}

// CreateProperty creates a new property
func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	var property models.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify landlord exists
	var landlord models.Landlord
	if err := h.DB.First(&landlord, property.LandlordID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Landlord not found"})
		return
	}

	if err := h.DB.Create(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property"})
		return
	}

	c.JSON(http.StatusCreated, property)
}

// GetProperties retrieves all properties with optional landlord info
func (h *PropertyHandler) GetProperties(c *gin.Context) {
	var properties []models.Property
	if err := h.DB.Preload("Landlord").Preload("ValuationMetrics").Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch properties"})
		return
	}

	c.JSON(http.StatusOK, properties)
}

// GetProperty retrieves a single property by ID
func (h *PropertyHandler) GetProperty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var property models.Property
	if err := h.DB.Preload("Landlord").Preload("ValuationMetrics").First(&property, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch property"})
		return
	}

	c.JSON(http.StatusOK, property)
}

// UpdateProperty updates a property
func (h *PropertyHandler) UpdateProperty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var property models.Property
	if err := h.DB.First(&property, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch property"})
		return
	}

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update property"})
		return
	}

	c.JSON(http.StatusOK, property)
}

// DeleteProperty deletes a property
func (h *PropertyHandler) DeleteProperty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.Delete(&models.Property{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete property"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Property deleted"})
}