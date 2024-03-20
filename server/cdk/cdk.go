package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// timestamp needed because bucket name has to be globally unique
	timestamp := time.Now().Unix()
	timestampString := strconv.Itoa(int(timestamp))
	awss3.NewBucket(stack, jsii.String("scribe-videos-"+timestampString), &awss3.BucketProps{})

	snsTopic := awssns.NewTopic(stack, jsii.String("scribe-topic"), &awssns.TopicProps{
		DisplayName:               jsii.String("scribe-topic"),
		ContentBasedDeduplication: jsii.Bool(true),
		Fifo:                      jsii.Bool(true),
		TopicName:                 jsii.String("scribe-topic"),
	})

	sqsQueue := awssqs.NewQueue(stack, jsii.String("scribe-audio-to-process"), &awssqs.QueueProps{
		QueueName:                 jsii.String("scribe-audio-to-process.fifo"),
		ContentBasedDeduplication: jsii.Bool(true),
		Fifo:                      jsii.Bool(true),
	})

	snsTopic.AddSubscription(awssnssubscriptions.NewSqsSubscription(sqsQueue, &awssnssubscriptions.SqsSubscriptionProps{}))

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "scribe-stack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	account := os.Getenv("ACCOUNT")
	region := os.Getenv("REGION")
	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
