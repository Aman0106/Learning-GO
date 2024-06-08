package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Aman:aman@golearningcluster0.vypppd8.mongodb.net/?retryWrites=true&w=majority&appName=GoLearningCluster0"
const dbName = "netflix"
const collectionName = "watchList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	handleError(err)
	fmt.Println("Successfully connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection refrence is set")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
