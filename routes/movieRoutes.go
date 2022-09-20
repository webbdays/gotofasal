package routes

import (
	"github.com/gorilla/mux"
	"github.com/webbdays/gotofasal/controllers"
)

func main() {

	var MovieRoutes = func(router *mux.Router) {
		router.HandleFunc("/movie", controllers.GetMovie).Methods("GET")
		router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")
		router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
		router.HandleFunc("/movie", controllers.UpdateMovie).Methods("PUT")
		router.HandleFunc("/movie", controllers.DeleteMovie).Methods("DELETE")
	}
}
