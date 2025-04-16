package main

import (
	"github.com/keylab/celestialbound/backend/routes"
)

func main() {
	router := routes.SetupRoutes()
	router.Run(":8080") // Starts the server on localhost:8080
}
