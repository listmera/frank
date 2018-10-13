package models

import (
	"github.com/listmera/frank/structs"
	"github.com/mongodb/mongo-go-driver/bson"
)

func InsertUser (user structs.ListmeraUser) {
	bson.NewDocument(
		bson.EC.String("birthdate", user.Birthdate),
		bson.EC.String("country", user.Country),
		bson.EC.String("username", user.UserName),
		bson.EC.String("href", user.Href),
		bson.EC.String("spotify_id", user.UserId),
		bson.EC.String("profile_img", user.ProfileImg),
		bson.EC.Int64("followers", user.Followers),
		)
}
