package mDB

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/notes_db"
const db = "events"
const collName = "events"

var mongoClient *mongo.Client

func ConnectMongoDB() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	mongoClient = client
}
