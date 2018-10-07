package controllers

import (
	"github.com/naoina/denco"
	"net/http"
)

type check struct {
	Message string `json:"message"`
}

func Base (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
}