package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func Publish(message string, topicArn string) string {
	config := LoadConfig()
	client := sns.NewFromConfig(config)

	result, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message: &message,
		TopicArn: &topicArn,
	})

	if err != nil {
		fmt.Println("Error publishing message: ", err)
	}

	return *result.MessageId
}