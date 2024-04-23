package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

func Publish(message string, topicArn string) string {
	config := LoadConfig()
	client := sns.NewFromConfig(config)

    attributes := map[string] types.MessageAttributeValue {
        "operation": {
			DataType: aws.String("String"),
			StringValue: aws.String("EXTRACT_AUDIO_TO_S3"),
		},
    }

	result, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message: &message,
		TopicArn: &topicArn,
		MessageAttributes: attributes,
	})

	if err != nil {
		fmt.Println("Error publishing message: ", err)
	}

	return *result.MessageId
}