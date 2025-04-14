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

func GetPlayerHandler(c *gin.Context) {

	//Get player ID from the request
	playerID := c.Param("player_id")

	//Check if the player exists
	playerState, exists := playerStates[playerID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Return the player state
	c.JSON(http.StatusOK, playerState)
}

func GetAllPlayersHandler(c *gin.Context) {
	var players []*models.PlayerState
	for _, player := range playerStates {
		players = append(players, player)
	}
	c.JSON(http.StatusOK, players)
}

func DeletePlayerHandler(c *gin.Context) {
	//Get player ID from the request
	playerID := c.Param("player_id")

	//Get player name linked to the ID
	playerName := playerStates[playerID].PlayerName

	//Check if the player exists
	_, exists := playerStates[playerID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Delete the player state
	delete(playerStates, playerID)

	// Return success message with player name
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully", "player_name": playerName, "player_id": playerID})
}

func UpdatePlayerHandler(c *gin.Context) {
	//Get player ID from the request
	playerID := c.Param("player_id")

	//Check if the player exists
	playerState, exists := playerStates[playerID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Parse incoming JSON request
	var input struct {
		PlayerName string `json:"player_name"`
		Stars      int    `json:"stars" `
		JarLevel   int    `json:"jar_level"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update player state
	if input.PlayerName != "" { // Allow change of name
		playerState.PlayerName = input.PlayerName
	}
	if input.Stars >= 0 { // Allow stars reset
		playerState.Stars = input.Stars
	}
	if input.JarLevel > 0 { // Allow jar level reset (base Jar level is 1)
		playerState.JarLevel = input.JarLevel
	}

	// Return the updated player state
	c.JSON(http.StatusOK, playerState)
}
