package handlers

import (
	"modernPokedex/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RoutesHandler(db *database.Postgres, rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	h := &Handler{
		db:  db,
		rdb: rdb,
	}

	r.GET("/", h.getAllPokemons)
	r.GET("/:id", h.getPokemon)

	return r
}
