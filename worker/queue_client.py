import os
import boto3
from dotenv import load_dotenv

load_dotenv()

sqs_client = boto3.client(
    'sqs',
    aws_access_key_id=os.getenv('AWS_KEY'),
    aws_secret_access_key=os.getenv('AWS_SECRET'),
)

QUEUE_URL = os.getenv('AWS_SQS_URL')