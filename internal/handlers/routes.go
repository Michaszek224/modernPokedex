package handlers

import "github.com/gin-gonic/gin"

func RoutesHandler() *gin.Engine {
	r := gin.Default()
	r.GET("/", getAllPokemons)
	r.GET("/:id", getPokemonByID)
	r.POST("/:id", createPokemon)
	return r
}
