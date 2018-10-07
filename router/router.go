package router

import (
	"github.com/naoina/denco"
	"net/http"
	"go-server/controllers"
)

func NewRouter() http.Handler {
	mux := denco.NewMux()

	router, err := mux.Build([]denco.Handler{
		mux.GET("/", controllers.Base),
		mux.GET("/check", controllers.Check),
	})

	if err != nil {
		panic(err)
	}
	return router
}