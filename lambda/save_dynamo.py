"""
save_dynamo.py
transcribeが保存したjsonファイルを読み込んで、字幕データをdynamodbに保存する
"""

import os
import json
import boto3


def lambda_handler(event, context):
    try:
        bucket_name = os.environ("BUCKET_NAME")
        id = event["detail"]["TranscriptionJobName"]

        s3 = boto3.client("s3")
        key = f"{id}.json"
        response = s3.get_object(Bucket=bucket_name, Key=key)
        content = response["Body"].read().decode("utf-8")
        json_dict = json.loads(content)

        dynamodb = boto3.resource("dynamodb", region_name="us-west-2")
        table_name = "transcribe"
        table = dynamodb.Table(table_name)
        invalid_data = ""
        with table.batch_writer() as batch:
            for item in json_dict["results"]["items"]:
                if (
                    item.get("start_time", None) is None
                    and item.get("end_time", None) is None
                ):
                    # start_timeとend_timeが空の場合は句読点
                    invalid_data = item["alternatives"][0]["content"]
                    continue
                if invalid_data != "":
                    # 1つ前のデータが句読点の場合は先頭に付け加える
                    item["alternatives"][0]["content"] = (
                        invalid_data + item["alternatives"][0]["content"]
                    )
                invalid_data = ""
                item["media_id"] = id
                batch.put_item(
                    Item={
                        "id": item["id"],
                        "media_id": str(item["media_id"]),
                        "type": item["type"],
                        "alternatives": item["alternatives"],
                        "start_time": item["start_time"],
                        "end_time": item["end_time"],
                    }
                )
    except Exception as e:
        print(str(e))
    return
