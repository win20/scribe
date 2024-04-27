from dataclasses import dataclass
from dotenv import load_dotenv
import os

load_dotenv()

@dataclass
class config:
    aws_key: str = os.getenv('AWS_KEY')
    aws_secret: str = os.getenv('AWS_SECRET')
    aws_sqs_url: str = os.getenv('AWS_SQS_URL')
    aws_bucket_name: str = os.getenv('AWS_BUCKET_NAME')
