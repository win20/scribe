package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "pong")
	})

	router.Run("localhost:8080")
}
