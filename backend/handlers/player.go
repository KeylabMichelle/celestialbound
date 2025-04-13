package handlers

import (
	"net/http" // for HTTP status codes

	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // for generating unique player IDs
	"github.com/keylab/celestialbound/backend/models"
)

func CreatePlayerHandler(c *gin.Context) {
	// Parse incoming JSON request
	var input struct {
		PlayerName string `json:"player_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate a unique PlayerID
	playerID := uuid.New().String()

	// Default jar
	starterJar := models.Jar{
		JarLevel:       1,
		StarsStored:    0,
		StarsPerSecond: 2,
		UpgradeCost:    100,
		MaxCapacity:    500,
	}

	// Initialize player state
	newPlayer := &models.PlayerState{
		PlayerID:       playerID,
		PlayerName:     input.PlayerName,
		Stars:          0,
		JarLevel:       1,
		StarsPerClick:  1,
		StarsPerSecond: 2,
		UpgradeCost:    100,
		Jars:           []models.Jar{starterJar},
		// -- Boost mechanics (for future use) ---
		/* ClickMultiplier:     1.0,
		ClickBoostExpiresAt: time.Time{},
		PassiveMultiplier:     1.0,
		PassiveBoostExpiresAt: time.Time{}, */
	}

	// Store the new player state
	playerStates[playerID] = newPlayer // This changes later when there is a DB

	// Return the new player state
	c.JSON(http.StatusCreated, newPlayer)
}
