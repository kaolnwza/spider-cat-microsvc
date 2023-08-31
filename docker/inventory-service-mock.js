db = db.getSiblingDB('inventory-service')

db.Item.insertMany(
    {
        //damage = 47
        "uuid": "f226518e-9096-4576-b229-03af17b78594",
        "name": "ak47",
    },
    {
        "uuid": "ce716411-e05e-43db-bbd5-6649c294c572",
        "name": "ak47_bullet"
    },
    {
         //healing = 47
        "uuid": "2a7ff7a3-d828-4973-8c8e-63d6bb704dea",
        "name": "tango",
    },
    {
        "uuid": "5d8e6f70-3a22-4476-b9d2-8014f8dd949e",
        "name": "spider_ass",
    },
    {
        "uuid": "f2888105-8216-4c37-8f50-37f080ac8d6c",
        "name": "guitar",
    },
    {
        "uuid": "ff03cc95-462b-408c-a46d-9f97b07b74ab",
        "name": "samurai",
    },
    {
        "uuid": "c9e5c3d8-c1a4-47ca-b086-fa9767280158",
        "name": "cheat_engine",
    }
)

db.Inventory.insertMany(
    {
        "item_uuid": "f226518e-9096-4576-b229-03af17b78594",
        "amount": 1
    }
)