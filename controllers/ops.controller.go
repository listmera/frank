package controllers

import (
	"encoding/json"
	"github.com/listmera/frank/utils"
	"github.com/naoina/denco"
	"net/http"
)

type check struct {
	Message string `json:"message"`
}

func Base (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func Check (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	result := check{"I've been: a puppet, a pauper, a pirate, a poet... A pawn and a king!"}
	err := json.NewEncoder(w).Encode(result)
	utils.CheckErr(err)
}