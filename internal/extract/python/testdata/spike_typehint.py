from typing import Any, Optional
import boto3

def get_item(
    table_name: str,
    key_id: str,
    default_val: Optional[dict[str, Any]] = None
) -> dict[str, Any]:
    dynamodb = boto3.resource("dynamodb")
    table = dynamodb.Table(table_name)
    response = table.get_item(Key={"id": key_id})
    return response.get("Item", default_val or {})

