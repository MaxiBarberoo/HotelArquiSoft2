package db

import (
	"fichadehotel/clients/hotel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var db *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Failed to ping the MongoDB server: ", err)
	}

	db = client.Database("prueba")

	// Set up the database for the hotel client
	hotel.Db = db
}

func StartDbEngine() {
	// We need to create or migrate collections
	hotelsCollection := db.Collection("hotels")

	// Define any indexes if needed
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"_id", 1}}, // Index on "_id" in ascending order
	}
	_, err := hotelsCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal("Failed to create an index: ", err)
	}

	log.Println("Finished Creating/Migrating Collections")
}
