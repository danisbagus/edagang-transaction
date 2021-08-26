package domain

import "time"

type TransactionModel struct {
	TransactionID   string    `bson:"transaction_id"`
	UserID          int64     `bson:"user_id"`
	ProductID       int64     `bson:"product_id"`
	Quantity        int64     `bson:"quantity"`
	TransactionDate time.Time `bson:"transaction_date"`
}
