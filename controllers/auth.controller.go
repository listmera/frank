package controllers

import (
	"github.com/naoina/denco"
	"net/http"
)

func Login (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func Register (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}