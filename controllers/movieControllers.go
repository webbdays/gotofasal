package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type movie struct{
	Name string
	Description string
	Director string
	Actors []string
	HeroName string
	HeroinName string
	YearOfRelease string
}


type person struct{
	FirstName string
	LastName string
	Age int
	Roles []string
}



func GetMovie(w http.ResponseWriter , request *http.Request){
	//aMovie []*movie;
	// logic to get movie by ID from db and save movie information in aMovie variable


	// json.NewEncoder(w).Encode(aMovie)
	return
	w.send("There are the movies")
}

func GetMovieByID(w http.ResponseWriter , request *http.Request){
	movieId = mux.Vars().movieId
	aMovie *movie;
	// logic to get movie by ID from db and save movie information in aMovie variable


	json.NewEncoder(w).Encode(aMovie)
	return
}

