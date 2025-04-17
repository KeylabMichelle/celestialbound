package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/handlers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CelestialBound Online",
		})
	})

	// Player routes
	playerGroup := router.Group("/player")
	{
		playerGroup.POST("/", handlers.CreatePlayerHandler)
		playerGroup.GET("/", handlers.GetAllPlayersHandler)
		playerGroup.GET("/:player_id", handlers.GetPlayerHandler)
		playerGroup.PUT("/:player_id", handlers.UpdatePlayerNameHandler)
		playerGroup.DELETE("/:player_id", handlers.DeletePlayerHandler)
	}

	// Jar routes
	jarsGroup := router.Group("/player/:player_id/jar")
	{
		jarsGroup.POST("/", handlers.CreateNewJarHandler)
		jarsGroup.GET("/", handlers.GetAllJarsHandler)
	}

	// Click route
	router.POST("/click/:player_id", handlers.ClickHandler)

	return router
}
