db = db.getSiblingDB('inventory-service');

db.Item.insertOne(
   {
      "hp": 1000
   }
)