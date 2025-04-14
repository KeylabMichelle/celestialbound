package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/handlers" // Importing handlers for player and click actions
)

func main() {
	router := gin.Default()

	// Define routes
	router.POST("/player", handlers.CreatePlayerHandler)        // Endpoint to create a new player
	router.POST("/click/:player_id", handlers.ClickHandler)     // Endpoint to handle clicks
	router.GET("/player/:player_id", handlers.GetPlayerHandler) // Endpoint to get player state

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CelestialBound Online",
		})
	})

	router.Run(":8080") // Starts the server on localhost:8080
}
