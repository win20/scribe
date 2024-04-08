package handlers

import (
	"scribe/server/helpers"
	"scribe/server/services"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Url string `json:"url"`
}

/* When we hit this endpoint, send url to aws...
   url goes to queue, lambda worker is notified and extracts the audio from video...
   worker on server picks up audio and transcribes it */
func InitiateTranscription(c *fiber.Ctx) error {
	topicArn := helpers.GetDotenv().ScribeTopicArn
	url := c.Query("url")

	object := Message{
		Url: url,
	}

	messageString := helpers.ObjectToString(object)
	messageId := services.Publish(messageString, topicArn)
	return c.Status(fiber.StatusOK).SendString(messageId)
}

func extractAudioFromVideoLink(url string) {

}

func storeVideo() {

}
