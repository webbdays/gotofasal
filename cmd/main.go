package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/webbdays/gotofasal/pkg/connections"
	"github.com/webbdays/gotofasal/pkg/routes"
)

func main() {
	// connect to a database cluster
	connections.ConnectToMongoDB()

	// define a root router
	router := mux.NewRouter()

	// switch on the movie routes
	routes.MovieRoutes(router)

	// listen and serve the application
	func() {
		fmt.Println("Server is starting to run on port 3000")
		log.Fatal(http.ListenAndServe(":3000", router))
	}()

}
