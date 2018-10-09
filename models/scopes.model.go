package models

type scopes []string

var ListmeraScopes = scopes{
	"user-read-private",
	"user-read-email",
	"user-read-recently-played",
	"user-library-read",
	"user-read-birthdate",
	"user-top-read",
	"playlist-read-private",
	"playlist-read-collaborative",
	"playlist-modify-public",
	"playlist-modify-private",
}