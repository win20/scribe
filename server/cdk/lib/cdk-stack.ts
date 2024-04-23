import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import { Topic, SubscriptionFilter } from 'aws-cdk-lib/aws-sns';
import { SqsSubscription } from 'aws-cdk-lib/aws-sns-subscriptions';
import { Queue } from 'aws-cdk-lib/aws-sqs';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { Runtime, Function, Code } from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
import { Construct } from 'constructs';
import * as path from 'path';
import { getConfig } from './config';
import { PolicyStatement } from 'aws-cdk-lib/aws-iam';

const config = getConfig();

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const queue = new Queue(this, `${config.APP_NAME}-queue`, {
      queueName: `${config.APP_NAME}-extract-audio-to-s3`,
      visibilityTimeout: Duration.seconds(300)
    });

    const queueAudioToBeTranscribed = new Queue(this, `${config.APP_NAME}-audio-to-be-transcribed`, {
      queueName: `${config.APP_NAME}-audio-to-be-transcribed`,
      visibilityTimeout: Duration.seconds(300)
    })

    const topic = new Topic(this, `${config.APP_NAME}-topic`, {
      topicName: config.APP_NAME,
      displayName: config.APP_NAME,
    });

    const lambdaExtractAudioToS3 = new Function(this, 'scribe-extract-audio-to-s3', {
      functionName: `${config.APP_NAME}-extract-audio-to-s3`,
      runtime: Runtime.NODEJS_20_X,
      handler: 'index.handler',
      code: Code.fromAsset(path.join(__dirname, '../lambda/extractAudioToS3')),
      timeout: Duration.seconds(20),
    });

    const bucket = new Bucket(this, config.APP_NAME, {
      bucketName: `${config.APP_NAME}-${config.ACCOUNT_NUMBER}-${config.REGION}`
    });

    const api = new RestApi(this, `${config.APP_NAME}-api`, {
      restApiName: config.APP_NAME
    });

    const extractAudioToS3Integration = new LambdaIntegration(lambdaExtractAudioToS3, {
      requestTemplates: {'application/json': '{ "statusCode": "200"'}
    });

    api.root.addMethod('POST', extractAudioToS3Integration);
    const snsPublishPolicy = new PolicyStatement({
      actions: ['sns:publish'],
      resources: ['*'],
    });
    lambdaExtractAudioToS3.addToRolePolicy(snsPublishPolicy)

    topic.addSubscription(new SqsSubscription(queue, {
      filterPolicy: {
        operation: SubscriptionFilter.stringFilter({
          allowlist: ['EXTRACT_AUDIO_TO_S3']
        })
      }
    }));
    topic.addSubscription(new SqsSubscription(queueAudioToBeTranscribed, {
      filterPolicy: {
        operation: SubscriptionFilter.stringFilter({
          allowlist: ['AUDIO_TO_BE_TRANSCRIBED']
        })
      }
    }));
    bucket.grantReadWrite(lambdaExtractAudioToS3);

    const eventSource = new SqsEventSource(queue)
    lambdaExtractAudioToS3.addEventSource(eventSource)
  }
}


