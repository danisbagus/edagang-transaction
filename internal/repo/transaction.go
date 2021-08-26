package repo

import (
	"context"

	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-pkg/logger"
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

func (r TransactionRepo) Create(data *domain.TransactionModel) *errs.AppError {
	collection := r.db.Database("edagang").Collection("transactions")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("Error on create transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}

func (r TransactionRepo) FindAll() ([]domain.TransactionModel, *errs.AppError) {
	filter := bson.M{}
	transactions := make([]domain.TransactionModel, 0)

	collection := r.db.Database("edagang").Collection("transactions")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		logger.Error("Error on Finding all transactions: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	for cur.Next(context.TODO()) {
		var transaction domain.TransactionModel
		err = cur.Decode(&transaction)
		if err != nil {
			logger.Error("Error on Decoding the transactionss: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r TransactionRepo) FindOneByID(transactionID string) (*domain.TransactionModel, *errs.AppError) {
	filter := bson.M{"transaction_id": transactionID}
	var transaction domain.TransactionModel

	collection := r.db.Database("edagang").Collection("transactions")
	err := collection.FindOne(context.TODO(), filter).Decode(&transaction)
	if err != nil {
		logger.Error("Error while one transaction " + err.Error())
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewNotFoundError("Transaction not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &transaction, nil
}

func (r TransactionRepo) Delete(transactionID string) *errs.AppError {
	filter := bson.M{"transaction_id": transactionID}

	collection := r.db.Database("edagang").Collection("transactions")
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		logger.Error("Error on delete transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
