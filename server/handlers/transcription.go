package handlers

import (
	"scribe/server/helpers"
	"scribe/server/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Message struct {
	JobUuid string `json:"jobUuid"`
	YoutubeUrl string `json:"youtubeUrl"`
}

/* When we hit this endpoint, send url to aws...
   url goes to queue, lambda worker is notified and extracts the audio from video...
   worker on server picks up audio and transcribes it */
func InitiateTranscription(c *fiber.Ctx) error {
	topicArn := helpers.GetDotenv().ScribeTopicArn
	youtubeUrl := c.Query("youtubeUrl")
	jobUuid := uuid.New()

	object := Message{
		JobUuid: jobUuid.String(),
		YoutubeUrl: youtubeUrl,
	}

	messageString := helpers.ObjectToString(object)
	messageId := services.Publish(messageString, topicArn)
	return c.Status(fiber.StatusOK).SendString(messageId)
}
