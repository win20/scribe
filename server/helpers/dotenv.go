package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

type Dotenv struct {
	ScribeTopicArn string
	ScribeQueueArn string
}

func GetDotenv() *Dotenv {
	godotenv.Load()

	return &Dotenv{
		ScribeTopicArn: os.Getenv("SCRIBE_TOPIC_ARN"),
		ScribeQueueArn: os.Getenv("SCRIBE_QUEUE_ARN"),
	}
}