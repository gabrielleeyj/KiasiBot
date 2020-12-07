package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDB : helper function to connect to mongoDB database
func ConnectDB() *mongo.Collection {
	// // Set client connection
	// clientOptions := options.Client().ApplyURI("mongodb+srv://Kayige:Kayige26387@kiasibot.egzpq.mongodb.net/db?retryWrites=true&w=majority")

	// // Start connection
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Kayige:Kayige26387@kiasibot.egzpq.mongodb.net/db?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// ping cluster to check connection status
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// check database connection status
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	fmt.Println("Connected to MongoDB Database!")

	collection := client.Database("db").Collection("usr")
	return collection
}

// ErrorResponse : Error Model Response
type ErrorResponse struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}

// GetError helper function that uses the Error Model Response
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
