db = db.getSiblingDB("edagang");

db.createCollection("transactions");

db.transactions.insertMany([
  {
    transaction_id: "TRGK8R6O",
    user_id: 1,
    product_id: 1,
    quantity: 1,
    transaction_date: new Date(Date.now()).toISOString(),
  },
  {
    transaction_id: "TR908I64",
    user_id: 2,
    product_id: 1,
    quantity: 2,
    transaction_date: new Date(Date.now()).toISOString(),
  },
  {
    transaction_id: "TR7HTVJV",
    user_id: 1,
    product_id: 2,
    quantity: 2,
    transaction_date: new Date(Date.now()).toISOString(),
  },
]);
