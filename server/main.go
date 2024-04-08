package main

import (
	"scribe/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Post("/initiate-transcription", handlers.InitiateTranscription)

	app.Listen(":8000")
}
