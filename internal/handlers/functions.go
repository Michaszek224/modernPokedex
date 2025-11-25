package handlers

import (
	"modernPokedex/internal/database"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type Handler struct {
	db *database.Postgres
}

func (h *Handler) getAllPokemons(ctx *gin.Context) {
	pokemons, err := h.db.GetAllPokemon()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error getting pokemons",
		})
		return
	}
	ctx.JSON(http.StatusOK, pokemons)
}

func (h *Handler) getPokemon(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pokemon, err := h.db.GetPokemonByName(idStr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "pokemon not found",
			})
			return
		}
		ctx.JSON(http.StatusOK, pokemon)
		return
	}

	pokemon, err := h.db.GetPokemonByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "pokemon not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, pokemon)
}
