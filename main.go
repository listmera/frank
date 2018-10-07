package main

import (
	"frank/env"
)

func main() {
//	r := router.NewRouter()
//	handler := middleware.ApplyMiddleware(r)
//	err := http.ListenAndServe(":3000", handler)
//	if err != nil {
//		panic(err)
//	}
}

func init() {
	env.SetEnv()
}