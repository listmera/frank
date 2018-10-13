package models

import (
	"context"
	"github.com/listmera/frank/utils"
	"github.com/listmera/frank/env"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var users *mongo.Collection
var client *mongo.Client

func Connect () {
	client, err := mongo.NewClient(env.GetOr("MONGO_URL", "mongodb://localhost:27017"))
	client.Connect(context.Background())
	utils.CheckErr(err)

	collection := client.Database(env.GetOr("MONGO_DB", "frank")).Collection(env.GetOr("MONGO_COLLECTION", "users"))

	users = collection
}

func Disconnect () {
	client.Disconnect(context.Background())
}