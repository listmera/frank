package controllers

import (
	"encoding/json"
	"github.com/listmera/frank/models"
	"github.com/listmera/frank/structs"
	"github.com/listmera/frank/utils"
	"github.com/naoina/denco"
	"net/http"
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

	//TODO: we use upsert in mongo to create new user/update existing one
	user := structs.NewListmeraUser(spotifyUser)
	id, err := models.InsertUser(user, tokens)
	utils.CheckErr(err)

	user.Id = id.InsertedID

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	utils.CheckErr(err)
	//TODO: add in the response a flag saying if its a new user or an old one
}

func Redirect (w http.ResponseWriter, r *http.Request, params denco.Params) {
	resBody := models.GenRedirect()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resBody)
	utils.CheckErr(err)
}