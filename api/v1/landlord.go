package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/house-hunt-labs/hh-service-property/pkg/models"
)

type LandlordHandler struct {
	DB *gorm.DB
}

func NewLandlordHandler(db *gorm.DB) *LandlordHandler {
	return &LandlordHandler{DB: db}
}

// CreateLandlord creates a new landlord
func (h *LandlordHandler) CreateLandlord(c *gin.Context) {
	var landlord models.Landlord
	if err := c.ShouldBindJSON(&landlord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&landlord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create landlord"})
		return
	}

	c.JSON(http.StatusCreated, landlord)
}

// GetLandlords retrieves all landlords
func (h *LandlordHandler) GetLandlords(c *gin.Context) {
	var landlords []models.Landlord
	if err := h.DB.Preload("Properties").Find(&landlords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch landlords"})
		return
	}

	c.JSON(http.StatusOK, landlords)
}

// GetLandlord retrieves a single landlord by ID
func (h *LandlordHandler) GetLandlord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var landlord models.Landlord
	if err := h.DB.Preload("Properties").First(&landlord, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Landlord not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch landlord"})
		return
	}

	c.JSON(http.StatusOK, landlord)
}

// UpdateLandlord updates a landlord
func (h *LandlordHandler) UpdateLandlord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var landlord models.Landlord
	if err := h.DB.First(&landlord, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Landlord not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch landlord"})
		return
	}

	if err := c.ShouldBindJSON(&landlord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&landlord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update landlord"})
		return
	}

	c.JSON(http.StatusOK, landlord)
}

// DeleteLandlord deletes a landlord
func (h *LandlordHandler) DeleteLandlord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.Delete(&models.Landlord{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete landlord"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Landlord deleted"})
}