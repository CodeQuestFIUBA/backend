package database

import (
	"context"

	"codequest/src/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoUri = configs.EnvMongoURI()

// A global variable that will hold a reference to the MongoDB client
var MongoClient *mongo.Client

// Implementation logic for connecting to MongoDB
func CreateMongoDBInstance() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoUri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	MongoClient = client

	if err != nil {
		panic(err)
	}

	return err
}

// Retrieves a handle to a specific collection within the database that the client is connected to
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(configs.EnvDBName()).Collection(collectionName)

	return collection
}
