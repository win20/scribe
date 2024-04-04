import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import { Topic } from 'aws-cdk-lib/aws-sns';
import { SqsSubscription } from 'aws-cdk-lib/aws-sns-subscriptions';
import { Queue } from 'aws-cdk-lib/aws-sqs';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { Runtime, Function, Code } from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';
import * as path from 'path';
import { getConfig } from './config';

const config = getConfig();

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const queue = new Queue(this, `${config.APP_NAME}-queue`, {
      queueName: config.APP_NAME,
      visibilityTimeout: Duration.seconds(300)
    });

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
    topic.addSubscription(new SqsSubscription(queue));
    bucket.grantReadWrite(lambdaExtractAudioToS3);
  }
}


