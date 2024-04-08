package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig() aws.Config {
	config, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading aws sdk: ", err)
	}

	return config
}
