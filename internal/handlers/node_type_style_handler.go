package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"github.com/house-hunt-labs/hh-service-property/internal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NodeTypeStyleHandler struct {
	service *services.NodeTypeStyleService
}

func NewNodeTypeStyleHandler(service *services.NodeTypeStyleService) *NodeTypeStyleHandler {
	return &NodeTypeStyleHandler{service: service}
}

func (h *NodeTypeStyleHandler) Create(c *gin.Context) {
	var style models.NodeTypeStyle
	if err := c.ShouldBindJSON(&style); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	style.ID = primitive.NewObjectID()
	err := h.service.Create(c.Request.Context(), &style)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, style)
}

func (h *NodeTypeStyleHandler) GetAll(c *gin.Context) {
	// Check for ?type query parameter
	nodeType := c.Query("type")
	if nodeType != "" {
		// Get by type if query parameter is provided
		style, err := h.service.GetByType(c.Request.Context(), nodeType)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Style not found"})
			return
		}
		c.JSON(http.StatusOK, style)
		return
	}
	// Get all styles if no query parameter
	styles, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, styles)
}

func (h *NodeTypeStyleHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	style, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, style)
}

func (h *NodeTypeStyleHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var style models.NodeTypeStyle
	if err := c.ShouldBindJSON(&style); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Update(c.Request.Context(), id, &style)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, style)
}

func (h *NodeTypeStyleHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
