package structs

type redirectRes struct {
	Redirect string `json:"redirect"`
}

func NewRedirectRes (redirect string) redirectRes {
	return redirectRes{
		redirect,
	}
}

type RegisterReq struct {
	Code string `json:"code"`
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}