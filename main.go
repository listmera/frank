package main

import (
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/middleware"
	"github.com/listmera/frank/router"
	"github.com/listmera/frank/utils"
	"net/http"
)

func main() {
	r := router.NewRouter()
	handler := middleware.ApplyMiddleware(r)
	err := http.ListenAndServe(env.GetOr("FRANK_PORT", ":3000"), handler)
	utils.CheckErr(err)
}

func init() {
	env.SetEnv()
}