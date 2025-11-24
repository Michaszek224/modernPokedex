package main

import (
	"modernPokedex/internal/handlers"
)

func main() {
	router := handlers.RoutesHandler()
	if err := router.Run(); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
