package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"github.com/house-hunt-labs/hh-service-property/internal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapNodeHandler struct {
	service *services.MapNodeService
}

func NewMapNodeHandler(service *services.MapNodeService) *MapNodeHandler {
	return &MapNodeHandler{service: service}
}

func (h *MapNodeHandler) Create(c *gin.Context) {
	var node models.MapNode
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node.ID = primitive.NewObjectID()
	err := h.service.Create(c.Request.Context(), &node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, node)
}

func (h *MapNodeHandler) GetAll(c *gin.Context) {
	// Check for ?type query parameter
	nodeType := c.Query("type")
	if nodeType != "" {
		// Get by type if query parameter is provided
		nodes, err := h.service.GetByType(c.Request.Context(), nodeType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nodes)
		return
	}
	// Get all nodes if no query parameter
	nodes, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *MapNodeHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	node, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, node)
}

func (h *MapNodeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var node models.MapNode
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Update(c.Request.Context(), id, &node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

func (h *MapNodeHandler) Delete(c *gin.Context) {
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
