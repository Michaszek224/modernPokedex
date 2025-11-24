package handlers

import "github.com/gin-gonic/gin"

func RoutesHandler() *gin.Engine {
	r := gin.Default()
	r.GET("/", ping)
	return r
}
