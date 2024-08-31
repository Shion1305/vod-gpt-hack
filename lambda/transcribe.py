"""
transcribe.py
aws transcribeを実行する関数
"""

import os
import json
import boto3

bucket_name = "pirotyyy-test-bucket-mp4"


def lambda_handler(event, context):
    bucket_name = os.environ("BUCKET_NAME")
    output_bucket_name = os.environ("OUTPUT_BUCKET_NAME")
    body = json.loads(event["Records"][0]["body"])
    id = body["id"]
    file_name = body["fileName"]
    transcribe = boto3.client("transcribe")

    # 保存されるjsonファイル名を簡単に特定するため、idと同じにする
    job_name = id

    s3_uri = f"s3://{bucket_name}/{id}/{file_name}"

    response = transcribe.start_transcription_job(
        TranscriptionJobName=job_name,
        Media={"MediaFileUri": s3_uri},
        MediaFormat="mp4",
        LanguageCode="ja-JP",
        OutputBucketName=output_bucket_name,
    )

    # ジョブの情報を表示
    print(
        f"Started Transcription Job: {response['TranscriptionJob']['TranscriptionJobName']}"
    )
