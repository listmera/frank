package structs

import "github.com/mongodb/mongo-go-driver/mongo"

func NewRedirectRes (redirect string) RedirectRes {
	return RedirectRes{
		redirect,
	}
}

func NewListmeraUser (resUser SpotifyUser) ListmeraUser {
	return ListmeraUser{
		mongo.InsertOneResult{},
		resUser.Birthdate,
		resUser.Country,
		resUser.UserName,
		resUser.Email,
		resUser.Href,
		resUser.UserId,
		resUser.Images[0].Url,
		int64(resUser.Followers.Total),
		[]string{},
		[]string{},
	}
}