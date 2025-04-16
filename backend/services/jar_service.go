/* Jar logic */

package services

import (
	"net/http" // for HTTP status codes

	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/models"
	"github.com/keylab/celestialbound/backend/utils"
)

func CreateNewJarService(c *gin.Context) {
	// Extract PlayerID from the request
	playerID := c.Param("player_id")

	// Try to get a pointer to the actual player state
	playerState, exists := playerStates[playerID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Create a new jar with default values
	newJar := models.Jar{
		JarID:          utils.GenerateUniqueID(),
		JarLevel:       1,
		StarsStored:    0,
		StarsPerSecond: 2,
		UpgradeCost:    100,
		MaxCapacity:    500,
	}

	// Append the new jar to the player's jar slice by updating the map directly
	playerStates[playerID].Jars = append(playerState.Jars, newJar)

	// Return the newly added jar
	c.JSON(http.StatusOK, newJar)
}

func GetAllJarsService(c *gin.Context) {
	//Get player ID from the request
	playerID := c.Param("player_id")

	//Check if the player exists
	playerState, exists := playerStates[playerID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Return the jar state
	c.JSON(http.StatusOK, playerState.Jars)
}
