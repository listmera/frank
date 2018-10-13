package structs

func NewRedirectRes (redirect string) RedirectRes {
	return RedirectRes{
		redirect,
	}
}

func NewListmeraUser (resUser SpotifyUser) ListmeraUser {
	return ListmeraUser{
		resUser.Birthdate,
		resUser.Country,
		resUser.UserName,
		resUser.Email,
		resUser.Href,
		resUser.UserId,
		resUser.Images[0].Url,
		int64(resUser.Followers.Total),
	}
}