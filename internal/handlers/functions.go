package handlers

import "github.com/gin-gonic/gin"

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
