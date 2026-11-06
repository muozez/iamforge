import boto3

def handle_event(action_type: str, payload: dict):
    sqs = boto3.client("sqs")
    match action_type:
        case "order_created":
            sqs.send_message(QueueUrl="orders-queue", MessageBody="order")
        case "order_cancelled" | "order_refunded":
            sqs.send_message(QueueUrl="cancelled-queue", MessageBody="cancel")
        case _:
            pass

