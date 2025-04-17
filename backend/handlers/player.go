package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/services"
)

var playerService = &services.PlayerService{}

func CreatePlayerHandler(c *gin.Context) {
	var input struct {
		PlayerName string `json:"player_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"player_id": playerService.CreatePlayer(input.PlayerName)})

}

func GetPlayerHandler(c *gin.Context) {
	playerID := c.Param("player_id")
	player := playerService.GetPlayer(playerID)
	if player == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func GetAllPlayersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, playerService.GetAllPlayers())
}

func UpdatePlayerHandler(c *gin.Context) {
	/* Check if player exists */
	playerID := c.Param("player_id")
	player := playerService.GetPlayer(playerID)
	if player == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	/* If it exists parse the input and check for error */
	var input struct {
		PlayerName string `json:"player_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	/* Update the player name */
	playerService.UpdatePlayerName(playerID, input.PlayerName)
	c.JSON(http.StatusOK, gin.H{"message": "Player updated"})
}

func DeletePlayerHandler(c *gin.Context) {
	playerID := c.Param("player_id")
	player := playerService.GetPlayer(playerID)
	if player == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	playerService.DeletePlayer(playerID)

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
