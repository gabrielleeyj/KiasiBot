package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb database
func server() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Kayige:Kayige26387@kiasibot.egzpq.mongodb.net/db?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
}

func SendLoc(value interface{}) {
	loc := value
	collection := client.Database("db").Collection("location")
	insertResult, err := collection.InsertOne(context.TODO(), loc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Location: ", insertResult.InsertedID)

}

func GetLoc() string {
	// mongodb chart map
	var mapchart = "http://"
	return mapchart
}
