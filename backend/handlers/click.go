package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/keylab/celestialbound/backend/services"
)

func ClickHandler(c *gin.Context) {
	/* Service that adds stars on click */
	services.ClickService(c)
}
