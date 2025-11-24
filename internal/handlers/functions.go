package handlers

import "github.com/gin-gonic/gin"

func getAllPokemons(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get request all",
	})
}

func getPokemonByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get request",
	})
}

func createPokemon(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "post request",
	})
}
