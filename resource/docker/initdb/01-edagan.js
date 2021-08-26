db = db.getSiblingDB("edagang");

db.createCollection("transactions");

db.transactions.insertMany([
  {
    transaction_id: "TRGK8R6O",
    product_id: 1,
    quantity: 1,
    transaction_date: new Date(Date.now()).toISOString(),
  },
  {
    transaction_id: "TR908I64",
    product_id: 1,
    quantity: 2,
    transaction_date: new Date(Date.now()).toISOString(),
  },
  {
    transaction_id: "TR7HTVJV",
    product_id: 2,
    quantity: 2,
    transaction_date: new Date(Date.now()).toISOString(),
  },
]);
