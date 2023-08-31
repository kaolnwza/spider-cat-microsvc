db = db.getSiblingDB('spider-service')

db.Spider.insertMany(
    {
        "uuid": "3386a3c7-8417-4930-9c95-078cdfbb2ec1",
        "name": "Crystal Maiden Spider",
        "damage": 1,
        "hp": 1,
        "spider_amount": 100,
        "drop_item_uuid": "2a7ff7a3-d828-4973-8c8e-63d6bb704dea",
        "drop_item_amount": 1
    },
    {
        "uuid": "8fbfe806-4b77-4adb-93b3-86a262ddbd69",
        "name": "Sniper Spider",
        "damage": 47,
        "hp": 100,
        "spider_amount": 100,
        "drop_item_uuid": "ce716411-e05e-43db-bbd5-6649c294c572",
        "drop_item_amount": 10
    },
    {
        "uuid": "82dd2453-8ec9-4c2b-9b32-3b30500c395d",
        "name": "Singer Spider",
        "damage": 999,
        "hp": 48,
        "spider_amount": 1,
        "drop_item_uuid": "f2888105-8216-4c37-8f50-37f080ac8d6c",
        "drop_item_amount": 30
    },
    {
        "uuid": "ab2a10ba-6ac0-4c14-96f2-f82db5b4dc0c",
        "name": "Spider King",
        "damage": 500,
        "hp":500,
        "spider_amount": 1,
        "drop_item_uuid": "ff03cc95-462b-408c-a46d-9f97b07b74ab",
        "drop_item_amount": 1
    },
    {
        "uuid": "f51ff547-a90d-40fd-848e-011a941aad37",
        "name": "Darksoul Spider",
        "damage": 9999999,
        "hp": 9999999,
        "spider_amount": 0,
    }
)