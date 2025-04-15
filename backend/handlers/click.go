package handlers

import (
	"net/http" // for HTTP status codes

	"github.com/gin-gonic/gin"
)

func ClickHandler(c *gin.Context) {

	// Extract PlayerID from the request
	playerID := c.Param("player_id")

	// Check if the player exists
	playerState, exists := playerStates[playerID]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Update stars based on the click
	playerState.Stars += playerState.StarsPerClick

	// Update stars in jar (just jar 1 for now)
	jar := &playerState.Jars[0] // reference to the struct value
	// Update stars in the jar
	jar.StarsStored += playerState.StarsPerClick

	// TODO: Check if jar is full and handle overflow

	//Return update to player state
	c.JSON(http.StatusOK, playerState)
}
