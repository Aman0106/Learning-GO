package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Aman0106/ApiWithMongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Aman:aman@golearningcluster0.vypppd8.mongodb.net/?retryWrites=true&w=majority&appName=GoLearningCluster0"
const dbName = "netflix"
const collectionName = "watchList"

var collection *mongo.Collection

// Mongo DB Helpers Below

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	handleError(err)
	fmt.Println("Successfully connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection refrence is set")
}

func insertOneMovie(movie model.Netflix) {
	fmt.Println("Inserting One value into DB:", movie)
	inserted, err := collection.InsertOne(context.Background(), movie)
	handleError(err)

	fmt.Println("Inserted One value into DB with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	fmt.Println("Started update process for id: ", movieId)
	id, err := primitive.ObjectIDFromHex(movieId)
	handleError(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	handleError(err)

	fmt.Println("Updated movie with count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	fmt.Println("Started delete with id:", movieId)
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	handleError(err)
	fmt.Println("Items deleted:", deleteResult.DeletedCount)
}

func deleteAllMovies() int64 {
	fmt.Println("Deleting Everything")
	filter := bson.D{{}} // Pass nothing in it to select everything
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	handleError(err)
	fmt.Println("Number of Documents Deleted: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	fmt.Println("Getting all Movies")
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	defer cursor.Close(context.Background())
	handleError(err)

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
