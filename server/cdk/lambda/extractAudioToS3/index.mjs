import { S3Client } from "@aws-sdk/client-s3"
import { Upload } from "@aws-sdk/lib-storage";
import ytdl from "ytdl-core";

export const handler = async (event, context) => {
  for (const record of event.Records) {
    const bucketName = 'scribe-672047559645-eu-west-1'
    const messageBody = JSON.parse(record.body)
    const message = JSON.parse(messageBody.Message)

    const { youtubeUrl, filePath } = {
      youtubeUrl: message.youtubeUrl,
      filePath: message.filePath
    };

    // Download audio from YouTube URL
    const audioStream = ytdl(youtubeUrl, { filter: 'audioonly' });

    // Upload audio stream to S3
    const uploadParams = {
      Bucket: bucketName,
      Key: filePath,
      Body: audioStream
    };

    try {
      const upload = new Upload({
        client: new S3Client({}),
        params: uploadParams,
      })

      upload.on("httpUploadProgress", (progress) => {
        console.log(progress)
      })

      await upload.done()

      return {
        statusCode: 200,
        body: JSON.stringify({
          message: 'Audio uploaded successfully',
        })
      };
    } catch (error) {
      return {
        statusCode: 500,
        body: JSON.stringify({
          error: 'Error processing request',
          details: error.message
        })
      };
    }
  }
};