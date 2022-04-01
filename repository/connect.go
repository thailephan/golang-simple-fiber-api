package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx context.Context
	Client *mongo.Client
	FlashCardsCollection *mongo.Collection
)

func init() {
	var err error
	// Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	Ctx = context.TODO()	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary")
	Client, err = mongo.Connect(Ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(Ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	FlashCardsCollection = Client.Database("flashcards").Collection("flashcards")

	fmt.Println("Successfully connected and pinged")
}
