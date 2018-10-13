package models

import (
	"encoding/base64"
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/structs"
	"net/http"
	"net/url"
	"strings"
)

var listmeraScopes = [10]string{
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

func GenRedirect () structs.RedirectRes {
	scopes := strings.Join(listmeraScopes[:], " ")
	id := env.GetOr("SPOTIFY_ID", "test")
	uri := env.GetOr("SPOTIFY_REDIRECT_URI", "test2")

	encodedScopes := url.QueryEscape(scopes)
	encodedUri := url.QueryEscape(uri)

	redirectUrl := "https://accounts.spotify.com/authorize?response_type=code&client_id=" + id +
		"&scope=" + encodedScopes + "&redirect_uri=" + encodedUri
	return structs.NewRedirectRes(redirectUrl)
}

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