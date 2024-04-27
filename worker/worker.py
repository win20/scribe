from queue_client import sqs_client, QUEUE_URL


def worker():
    while True:
        response = sqs_client.receive_message(
            QueueUrl=QUEUE_URL,
            MaxNumberOfMessages=10,
            WaitTimeSeconds=10,
            MessageAttributeNames=['All'],
        )

        for message in response.get('Messages', []):
            print(message)

            sqs_client.delete_message(
                QueueUrl=QUEUE_URL,
                ReceiptHandle=message['ReceiptHandle']
            )
        else:
            print("No messages to process.")


def main():
    worker()


if __name__ == "__main__":
    main()