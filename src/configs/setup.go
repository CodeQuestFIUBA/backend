package configs

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoUri = EnvMongoURI()

// A global variable that will hold a reference to the MongoDB client
var MongoClient *mongo.Client

// Implementation logic for connecting to MongoDB
func ConnectToMongoDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoUri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	return err
}
