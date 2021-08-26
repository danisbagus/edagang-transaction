package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danisbagus/edagang-pkg/logger"
	"github.com/danisbagus/edagang-transaction/internal/core/service"
	"github.com/danisbagus/edagang-transaction/internal/handler"
	"github.com/danisbagus/edagang-transaction/internal/repo"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	client := GetClient()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	router := mux.NewRouter()

	transactionRepo := repo.NewTransactionRepo(client)
	TransactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.TransactionHandler{Service: TransactionService}

	// routing
	router.HandleFunc("/transactions", transactionHandler.GetTransactionList).Methods(http.MethodGet).Name("GetTransactionList")
	router.HandleFunc("/transactions/{transaction_id}", transactionHandler.GetTransactionDetail).Methods(http.MethodGet).Name("GetTransactionDetail")

	// router.HandleFunc("/transactions", transactionHandler.NewTransaction).Methods(http.MethodPost).Name("NewTransaction")

	// starting server
	logger.Info("Starting transaction service")
	log.Fatal(http.ListenAndServe("localhost:9020", router))
}

func GetClient() *mongo.Client {
	var cred options.Credential

	cred.AuthSource = "admin"
	cred.Username = "root"
	cred.Password = "danisbagus"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred) // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
