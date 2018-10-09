package router

import (
	"github.com/listmera/frank/controllers"
	"github.com/listmera/frank/utils"
	"github.com/naoina/denco"
	"net/http"
)

func NewRouter() http.Handler {
	mux := denco.NewMux()

	router, err := mux.Build([]denco.Handler{
		mux.GET("/", controllers.Base),
		mux.GET("/check", controllers.Check),
		mux.GET("/login", controllers.Login),
		mux.POST("/register", controllers.Register),
		mux.GET("/me", controllers.GetMe),
		mux.PUT("/me", controllers.EditMe),
		mux.GET("/me/refresh", controllers.RefreshToken),
		mux.GET("/user/:id", controllers.GetUser),
	})

	utils.CheckErr(err)
	return router
}