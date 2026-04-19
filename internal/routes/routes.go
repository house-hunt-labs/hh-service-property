package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/house-hunt-labs/hh-service-property/internal/handlers"
    "github.com/house-hunt-labs/hh-service-property/internal/repositories"
    "github.com/house-hunt-labs/hh-service-property/internal/services"
    "go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, db *mongo.Database) {
    // Repositories
    nodeTypeStyleRepo := repositories.NewNodeTypeStyleRepository(db)
    mapNodeRepo := repositories.NewMapNodeRepository(db)

    // Services
    nodeTypeStyleService := services.NewNodeTypeStyleService(nodeTypeStyleRepo)
    mapNodeService := services.NewMapNodeService(mapNodeRepo)

    // Handlers
    nodeTypeStyleHandler := handlers.NewNodeTypeStyleHandler(nodeTypeStyleService)
    mapNodeHandler := handlers.NewMapNodeHandler(mapNodeService)

    // Routes
    api := r.Group("/api")
    {
        styles := api.Group("/node-type-styles")
        {
            styles.POST("", nodeTypeStyleHandler.Create)
            styles.GET("", nodeTypeStyleHandler.GetAll)
            styles.GET("/:id", nodeTypeStyleHandler.GetByID)
            styles.PUT("/:id", nodeTypeStyleHandler.Update)
            styles.DELETE("/:id", nodeTypeStyleHandler.Delete)
        }

        nodes := api.Group("/map-nodes")
        {
            nodes.POST("", mapNodeHandler.Create)
            nodes.GET("", mapNodeHandler.GetAll)
            nodes.GET("/:id", mapNodeHandler.GetByID)
            nodes.PUT("/:id", mapNodeHandler.Update)
            nodes.DELETE("/:id", mapNodeHandler.Delete)
        }
    }
}