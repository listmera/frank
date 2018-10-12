package structs

func NewRedirectRes (redirect string) redirectRes {
	return redirectRes{
		redirect,
	}
}

func NewListmeraUser (resUser SpotifyUser) listmeraUser {
	return listmeraUser{
		resUser.Birthdate,
		resUser.Country,
		resUser.UserName,
		resUser.Email,
		resUser.Href,
		resUser.UserId,
		resUser.Images[0].Url,
		resUser.Followers.Total,
	}
}