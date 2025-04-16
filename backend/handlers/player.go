package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/services"
)

func CreatePlayerHandler(c *gin.Context) {
	services.CreatePlayerService(c)
}

func GetPlayerHandler(c *gin.Context) {
	services.GetPlayerService(c)
}

func GetAllPlayersHandler(c *gin.Context) {
	services.GetAllPlayersService(c)
}

func UpdatePlayerHandler(c *gin.Context) {
	services.UpdatePlayerService(c)
}

func DeletePlayerHandler(c *gin.Context) {
	services.DeletePlayerService(c)
}
