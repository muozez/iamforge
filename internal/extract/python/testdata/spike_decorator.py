from functools import wraps
import boto3

def retry_aws(max_attempts=3):
    def decorator(fn):
        @wraps(fn)
        def wrapper(*args, **kwargs):
            return fn(*args, **kwargs)
        return wrapper
    return decorator

@retry_aws(max_attempts=5)
def send_msg(queue_url, body):
    sqs = boto3.client("sqs")
    return sqs.send_message(QueueUrl=queue_url, MessageBody=body)

