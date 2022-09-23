package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type movie struct {
// 	Name string `json:"Name"`
// 	// Description   string
// 	// Director      string
// 	// Actors        []string
// 	// HeroName      string
// 	// HeroinName    string
// 	// YearOfRelease string
// }

// type person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Roles     []string
// }


func SearchMovies(w http.ResponseWriter, request *http.Request){
	
	return
}



func GetMovies(w http.ResponseWriter, request *http.Request) {
	//aMovie []*movie;
	// logic to get movie by ID from db and save movie information in aMovie variable

	// json.NewEncoder(w).Encode(aMovie)
	fmt.Fprintf(w, "movies\n")
	return
	// w.send("There are the movies")
}

func GetMovieById(w http.ResponseWriter, request *http.Request) {
	movieId := mux.Vars(request)["movieId"]
	fmt.Fprintf(w, "movie with id: %v\n", movieId)
	// aMovie * movie
	// logic to get movie by ID from db and save movie information in aMovie variable

	// json.NewEncoder(w).Encode(aMovie)
	return
}

func CreateMovie(w http.ResponseWriter, request *http.Request) {
	return
}

func DeleteMoviesByIds(w http.ResponseWriter, request *http.Request) {
	// get input as ids from the request body but we need to make sure to perform input validation.
	// var movieIds movie
	// _ = json.NewDecoder(request.Body).Decode(&movieIds)
	// fmt.Println(movieIds)
	// fmt.Fprintf(w, "deleted movies with id: %v\n", movieIds)
	return
}

func DeleteMovieById(w http.ResponseWriter, request *http.Request) {
	// get the id from the parameter
	movieId := mux.Vars(request)["movieId"]
	fmt.Fprintf(w, "deleted movie with id: %v\n", movieId)
	return
}

func UpdateMovie(w http.ResponseWriter, request *http.Request) {
	// get Updated Movie information from request body
	return
}
