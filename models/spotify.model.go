package models

import (
	"encoding/base64"
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/structs"
	"net/http"
	"net/url"
	"strings"
)

func GetTokens (code string) (*http.Response, error) {
	address := "https://accounts.spotify.com/api/token"
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", env.GetOr("SPOTIFY_REDIRECT_URI", "FAIL!"))

	auth := env.GetOr("SPOTIFY_ID", "FAIL") + ":" + env.GetOr("SPOTIFY_SECRET", "FAIL")
	b64Auth := "Basic: " + base64.StdEncoding.EncodeToString([]byte(auth))

	request, err := http.NewRequest("POST", address, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", b64Auth)

	client := &http.Client{}
	resp, err := client.Do(request)

	return resp, err
}

func GetMe (tokens structs.TokenRes) (*http.Response, error) {
	address := "https://api.spotify.com/v1/me"
	request, err := http.NewRequest("GET", address, nil)
	request.Header.Set("Authorization", tokens.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(request)

	return resp, err
}