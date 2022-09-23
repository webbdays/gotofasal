package routes

import (
	"github.com/gorilla/mux"
	"github.com/webbdays/gotofasal/pkg/controllers"
)

var LoginRoutes = func(router *mux.Router) {
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
}
