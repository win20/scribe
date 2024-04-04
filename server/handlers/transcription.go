package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transcription(c *gin.Context) {
	url := c.Query("url")
	c.String(http.StatusOK, "URL: %s", url)
}

func extractAudioFromVideoLink(url string) {

}

func storeVideo() {

}
