package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/handlers" // Importing handlers for player and click actions
)

func main() {
	router := gin.Default()

	/* Routes for player.go */
	router.POST("/player", handlers.CreatePlayerHandler)              // Endpoint to create a new player
	router.GET("/player/:player_id", handlers.GetPlayerHandler)       // Endpoint to get player state
	router.GET("/players", handlers.GetAllPlayersHandler)             // Endpoint to get all players
	router.DELETE("/player/:player_id", handlers.DeletePlayerHandler) // Endpoint to delete a player
	router.PUT("/player/:player_id", handlers.UpdatePlayerHandler)    // Endpoint to update player Name, Stars and JarLevel

	/* Jars on players */
	router.POST("/player/:player_id/jar", handlers.CreateNewJarHandler) // Endpoint to create a new jar for a player
	router.GET("/player/:player_id/jars", handlers.GetAllJarsHandler)   // Endpoint to get jar state

	/* Routes for click.go */
	router.POST("/click/:player_id", handlers.ClickHandler) // Endpoint to handle clicks (add stars)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CelestialBound Online",
		})
	})

	router.Run(":8080") // Starts the server on localhost:8080
}
