package services

import (
	"net/http" // for HTTP status codes

	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/models"
	"github.com/keylab/celestialbound/backend/utils"
)

// In-memory storage for player states
var playerStates = make(map[string]*models.PlayerState)

func CreatePlayerService(c *gin.Context) {
	// Parse incoming JSON request
	var input struct {
		PlayerName string `json:"player_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate a unique PlayerID
	playerID := utils.GenerateUniqueID()

	//Generate Jar unique ID
	jarID := utils.GenerateUniqueID()

	// Default jar
	starterJar := models.Jar{
		JarID:          jarID,
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

func GetPlayerService(c *gin.Context) {

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

func GetAllPlayersService(c *gin.Context) {
	var players []*models.PlayerState
	for _, player := range playerStates {
		players = append(players, player)
	}
	c.JSON(http.StatusOK, players)
}

func DeletePlayerService(c *gin.Context) {
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

func UpdatePlayerService(c *gin.Context) {
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

	// Return the updated player state
	c.JSON(http.StatusOK, playerState)
}
