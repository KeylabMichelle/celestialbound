package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/services"
)

var jarService = &services.JarService{}

func CreateNewJarHandler(c *gin.Context) {

	playerID := c.Param("player_id")

	err := jarService.CreateNewJar(playerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jar created successfully"})
}

func GetAllJarsHandler(c *gin.Context) {
	playerID := c.Param("player_id")

	jars, err := jarService.GetAllJars(playerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jars)
}
