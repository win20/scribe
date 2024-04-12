import { S3Client } from "@aws-sdk/client-s3"
import { Upload } from "@aws-sdk/lib-storage";
import ytdl from "ytdl-core";

const ROOT_FILE_PATH = 'audio/'
const FILE_TYPE_EXT = '.wav'

export const handler = async (event, context) => {
  for (const record of event.Records) {
    const bucketName = process.env.SCRIBE_BUCKET_NAME
    const messageBody = JSON.parse(record.body)
    const message = JSON.parse(messageBody.Message)

    const {jobUuid, youtubeUrl } = {
      jobUuid: message.jobUuid,
      youtubeUrl: message.youtubeUrl,
    };

    const audioStream = ytdl(youtubeUrl, { filter: 'audioonly' });

    const fullFilePath = `${ROOT_FILE_PATH}${jobUuid}${FILE_TYPE_EXT}`
    const uploadParams = {
      Bucket: bucketName,
      Key: fullFilePath,
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