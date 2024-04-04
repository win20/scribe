package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scribe/server/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "pong")
	})

	// router.GET("/transcription/:url", func(ctx *gin.Context) {
	// 	test := ctx.Param("url")
	// 	ctx.String(http.StatusOK, "URL: "+test)
	// })

	router.POST("/transcription", handlers.Transcription)

	router.Run("localhost:8080")
}
