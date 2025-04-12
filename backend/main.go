package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "CelestialBound Online",
        })
    })

    router.Run(":8080") // Starts the server on localhost:8080
}
