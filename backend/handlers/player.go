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

	newPlayer, err := playerService.CreatePlayer(input.PlayerName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"player_id": newPlayer})

}

func GetPlayerHandler(c *gin.Context) {
	playerID := c.Param("player_id")

	player, err := playerService.GetPlayer(playerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func GetAllPlayersHandler(c *gin.Context) {

	players, err := playerService.GetAllPlayers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)

}

func UpdatePlayerNameHandler(c *gin.Context) {
	playerID := c.Param("player_id")

	var input struct {
		PlayerName string `json:"player_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := playerService.UpdatePlayerName(playerID, input.PlayerName)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player updated"})
}

func DeletePlayerHandler(c *gin.Context) {
	playerID := c.Param("player_id")
	err := playerService.DeletePlayer(playerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
