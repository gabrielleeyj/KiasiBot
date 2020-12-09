package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

// ConnectDB : helper function to connect to mongoDB database
func ConnectDB() *mongo.Collection {
	// main code to start connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	// ping cluster to check connection status
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)

	// check available databases and prints output
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// No errors show success message
	fmt.Println("Connected to MongoDB Database!")

	// Connect to MongoDB collection to store data.
	collection := client.Database("db").Collection("usr")
	return collection
}
