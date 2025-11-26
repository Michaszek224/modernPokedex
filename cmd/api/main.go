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

	rdb, err := database.RedisInit()
	if err != nil {
		log.Fatalf("Error initializing redis: %v", err)
	}
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatalf("Error closing redis: %v", err)
		}
	}()

	postgres := database.NewPostgres(db)

	err = database.SeedData(db)
	if err != nil {
		log.Fatalf("Error populate db: %v", err)
	}

	router := handlers.RoutesHandler(postgres, rdb)

	if err := router.Run(); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
