package router

import (
	"github.com/listmera/frank/controllers"
	"github.com/naoina/denco"
	"net/http"
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