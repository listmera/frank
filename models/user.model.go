package models

import (
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/utils"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Connect () *mongo.Collection {
	client, err := mongo.NewClient(env.GetOr("MONGO_URL", "mongodb://localhost:27017"))
	utils.CheckErr(err)

	collection := client.Database(env.GetOr("MONGO_DB", "frank")).Collection(env.GetOr("MONGO_COLLECTION", "users"))

	return collection
}