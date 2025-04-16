package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/services"
)

func CreateNewJarHandler(c *gin.Context) {
	services.CreateNewJarService(c)
}

func GetAllJarsHandler(c *gin.Context) {
	services.GetAllJarsService(c)
}
