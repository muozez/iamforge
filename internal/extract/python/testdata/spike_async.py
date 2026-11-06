import asyncio
import aioboto3

async def upload_records(bucket_name: str, items: list[dict]):
    session = aioboto3.Session()
    async with session.client("s3") as s3:
        tasks = [
            s3.put_object(Bucket=bucket_name, Key=f"item_{i}.json", Body=str(item))
            for i, item in enumerate(items)
        ]
        await asyncio.gather(*tasks)

