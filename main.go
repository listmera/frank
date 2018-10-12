package main

import (
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/middleware"
	"github.com/listmera/frank/router"
	"github.com/listmera/frank/utils"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	handler := middleware.ApplyMiddleware(r)
	log.Printf("Frank running in %s%s", env.GetOr("FRANK_HOST", "localhost"), env.GetOr("FRANK_PORT", ":3000"))
	err := http.ListenAndServe(env.GetOr("FRANK_PORT", ":1212"), handler)
	utils.CheckErr(err)
}

func init() {
	env.SetEnv()
}