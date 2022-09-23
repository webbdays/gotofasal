package routes

import (
	"github.com/gorilla/mux"
	"github.com/webbdays/gotofasal/pkg/controllers"
)

var MovieRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.IndexPage).Methods("GET")
	router.HandleFunc("/search", controllers.SearchMovies).Methods("POST")
	router.HandleFunc("/movie", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie", controllers.DeleteMoviesByIds).Methods("DELETE")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovieById).Methods("DELETE")
}


