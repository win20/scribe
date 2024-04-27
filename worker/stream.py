import boto3
from pydub import AudioSegment
from pydub.playback import play
import io
from config import config


def stream_audio_from_s3(bucket_name, object_key):
    s3 = boto3.client('s3')

    response = s3.get_object(Bucket=bucket_name, Key=object_key)
    audio_data = response['Body'].read()

    audio_segment = AudioSegment.from_file(io.BytesIO(audio_data))
    play(audio_segment)


def main():
    audio_path = 'audio/efcc431a-9b4f-4fd3-b80b-c088bd8a8f87.wav'

    boto3.setup_default_session(aws_access_key_id=config.aws_key,
                                aws_secret_access_key=config.aws_secret)

    stream_audio_from_s3(config.aws_bucket_name, audio_path)


if __name__ == "__main__":
    main()