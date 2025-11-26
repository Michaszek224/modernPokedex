package handlers

import (
	"encoding/json"
	"fmt"
	"modernPokedex/internal/database"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

type Handler struct {
	db  *database.Postgres
	rdb *redis.Client
}

func (h *Handler) getAllPokemons(ctx *gin.Context) {
	cacheKey := "allPokemons"

	cached, err := h.rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		var pokemons []database.Pokemon
		if err := json.Unmarshal([]byte(cached), &pokemons); err == nil {
			ctx.Header("X-Cache", "HIT")
			ctx.JSON(http.StatusOK, pokemons)
			return
		}
	}

	pokemons, err := h.db.GetAllPokemon()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error getting pokemons",
		})
		return
	}

	pokemonsJSON, err := json.Marshal(pokemons)
	if err == nil {
		h.rdb.Set(ctx, cacheKey, pokemonsJSON, 5*time.Minute)
	}

	ctx.Header("X-Cache", "MISS")
	ctx.JSON(http.StatusOK, pokemons)
}

func (h *Handler) getPokemon(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		cachedKey := fmt.Sprintf("pokemon:name:%s", idStr)
		cachedData, err := h.rdb.Get(ctx, cachedKey).Result()
		if err == nil {
			var pokemon database.Pokemon
			if err := json.Unmarshal([]byte(cachedData), &pokemon); err == nil {
				ctx.Header("X-Cache", "HIT")
				ctx.JSON(http.StatusOK, pokemon)
				return
			}
		}

		pokemon, err := h.db.GetPokemonByName(idStr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "pokemon not found",
			})
			return
		}

		pokemonJSON, err := json.Marshal(pokemon)
		if err == nil {
			h.rdb.Set(ctx, cachedKey, pokemonJSON, 5*time.Minute)
		}
		ctx.Header("X-Cache", "MISS")
		ctx.JSON(http.StatusOK, pokemon)
		return
	}

	cachedKey := fmt.Sprintf("pokemon:id:%d", id)
	cachedData, err := h.rdb.Get(ctx, cachedKey).Result()
	if err == nil {
		var pokemon database.Pokemon
		if err := json.Unmarshal([]byte(cachedData), &pokemon); err == nil {
			ctx.Header("X-Cache", "HIT")
			ctx.JSON(http.StatusOK, pokemon)
			return
		}
	}

	pokemon, err := h.db.GetPokemonByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "pokemon not found",
		})
		return
	}

	pokemonJSON, err := json.Marshal(pokemon)
	if err == nil {
		h.rdb.Set(ctx, cachedKey, pokemonJSON, 5*time.Minute)
	}

	ctx.Header("X-Cache", "MISS")
	ctx.JSON(http.StatusOK, pokemon)
}
