package handlers

import (
	"modernPokedex/internal/database"

	"github.com/gin-gonic/gin"
)

func RoutesHandler(db *database.Postgres) *gin.Engine {
	r := gin.Default()

	h := &Handler{db: db}

	r.GET("/", h.getAllPokemons)
	r.GET("/:id", h.getPokemon)

	return r
}
