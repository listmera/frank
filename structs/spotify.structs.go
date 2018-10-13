package structs

type RedirectRes struct {
	Redirect string `json:"redirect"`
}

type RegisterReq struct {
	Code string `json:"code"`
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SpotifyUser struct {
	Birthdate string `json:"birthdate"`
	Country string `json:"country"`
	UserName string `json:"display_name"`
	Email string `json:"email"`
	Href string `json:"href"`
	UserId string `json:"id"`
	Images images `json:"images"`
	Followers followers `json:"followers"`
}

type followers struct {
	Total int `json:"total"`
}

type image struct {
	Url string `json:"url"`
}

type images []image

type ListmeraUser struct {
	Id string `json:"id"`
	Birthdate string `json:"birthdate"`
	Country string `json:"country"`
	UserName string `json:"display_name"`
	Email string `json:"email"`
	Href string `json:"href"`
	SpotifyId string `json:"id"`
	ProfileImg string `json:"profile_image"`
	Followers int64 `json:"followers"`
	Playlists []string `json:"playlists"`
	Tracks []string `json:"tracks"`
}