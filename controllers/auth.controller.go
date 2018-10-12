package controllers

import (
	"encoding/json"
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/models"
	"github.com/listmera/frank/structs"
	"github.com/listmera/frank/utils"
	"github.com/naoina/denco"
	"net/http"
	"net/url"
	"strings"
)

func Login (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func Register (w http.ResponseWriter, r *http.Request, params denco.Params) {
	decoder := json.NewDecoder(r.Body)

	var reqBody structs.RegisterReq
	err := decoder.Decode(&reqBody)
	utils.CheckErr(err)

	spotifyRes, err := models.GetTokens(reqBody.Code)
	defer spotifyRes.Body.Close()
	utils.CheckErr(err)

	var tokens structs.TokenRes
	decoder = json.NewDecoder(spotifyRes.Body)
	err = decoder.Decode(&tokens)
	utils.CheckErr(err)

	userRes, err := models.GetMe(tokens)
	defer userRes.Body.Close()
	utils.CheckErr(err)

	var spotifyUser structs.SpotifyUser
	decoder = json.NewDecoder(userRes.Body)
	err = decoder.Decode(&spotifyUser)
	utils.CheckErr(err)

	user := structs.NewListmeraUser(spotifyUser)
	//models.GetTokens(req.Code) // from here on, everything should be a goroutine
}

func Redirect (w http.ResponseWriter, r *http.Request, params denco.Params) {
	scopes := strings.Join(models.ListmeraScopes, " ")
	id := env.GetOr("SPOTIFY_ID", "test")
	uri := env.GetOr("SPOTIFY_REDIRECT_URI", "test2")

	encodedScopes := url.QueryEscape(scopes)
	encodedUri := url.QueryEscape(uri)

	redirectUrl := "https://accounts.spotify.com/authorize?response_type=code&client_id=" + id +
		"&scope=" + encodedScopes + "&redirect_uri=" + encodedUri
	resBody := structs.NewRedirectRes(redirectUrl)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resBody)
	utils.CheckErr(err)
}