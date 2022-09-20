package main

import (
	"github.com/gorilla/mux"
	"github.com/webbdays/gotofasal/routes"
)

func main() {
	router := mux.NewRouter()
	routes.MovieRoutes(r)

}
