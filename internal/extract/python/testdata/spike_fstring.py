import boto3

bucket = "my-bucket"
key = "data/file.json"
arn = f"arn:aws:s3:::{bucket}/{key}"

s3 = boto3.client("s3")
s3.put_object(Bucket=bucket, Key=f"{key}-v2")

