package controllers

import (
	"encoding/json"
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/models"
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
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

type redirectRes struct {
	Redirect string `json:"redirect"`
}

func Redirect (w http.ResponseWriter, r *http.Request, params denco.Params) {
	scopes := strings.Join(models.ListmeraScopes, " ")
	id := env.GetOr("SPOTIFY_ID", "test")
	uri := env.GetOr("SPOTIFY_REDIRECT_URI", "test2")

	redirect := redirectRes{
		"https://accounts.spotify.com/authorize?response_type=code&client_id=" + id + "&scope=" +
			url.QueryEscape(scopes) + "&redirect_uri=" + url.QueryEscape(uri),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(redirect)
	utils.CheckErr(err)
}