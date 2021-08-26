package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/domain"
	"github.com/danisbagus/edagang-transaction/internal/core/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	db *mongo.Client
}

func NewTransactionRepo(db *mongo.Client) port.ITransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (r TransactionRepo) FindAll() ([]domain.TransactionModel, *errs.AppError) {
	filter := bson.M{}
	transactions := make([]domain.TransactionModel, 0)

	trnsactionCollection := r.db.Database("edagangan").Collection("transactions")
	cur, err := trnsactionCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all transaction", err)
	}
	for cur.Next(context.TODO()) {
		var transaction domain.TransactionModel
		err = cur.Decode(&transaction)
		if err != nil {
			log.Fatal("Error on Decoding the transaction", err)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println("transactions", transactions)
	return transactions, nil
}
