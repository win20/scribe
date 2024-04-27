from config import config
import boto3

sqs_client = boto3.client(
    'sqs',
    aws_access_key_id=config.aws_key,
    aws_secret_access_key=config.aws_secret,
)


def worker():
    queue_url = config.aws_sqs_url

    while True:
        response = sqs_client.receive_message(
            QueueUrl=queue_url,
            MaxNumberOfMessages=10,
            WaitTimeSeconds=10,
            MessageAttributeNames=['All'],
        )

        for message in response.get('Messages', []):
            print(message)

            sqs_client.delete_message(
                QueueUrl=queue_url,
                ReceiptHandle=message['ReceiptHandle']
            )
        else:
            print("No messages to process.")


def main():
    worker()


if __name__ == "__main__":
    main()