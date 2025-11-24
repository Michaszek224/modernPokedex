package main

import (
	"modernPokedex/internal/handlers"
)

func main() {
	router := handlers.RoutesHandler()
	router.Run(":8080")
}
