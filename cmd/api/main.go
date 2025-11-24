package main

import (
	"log"
	"modernPokedex/internal/database"
	"modernPokedex/internal/handlers"
)

func main() {
	db, err := database.PostgresInit()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}()

	router := handlers.RoutesHandler()
	if err := router.Run(); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
