package controllers

// import (
// 	"net/http"

// 	"github.com/webbdays/gotofasal/pkg/models"
// 	"github.com/webbdays/gotofasal/templates"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func IndexPage(w http.ResponseWriter, request *http.Request) {
// 	t := templates.IndexPage()
// 	var movies []*models.Movie;
	
// 	movies = append(movies, &models.Movie{
// 		Id: primitive.ObjectID{},
// 		Name: "movie1",
// 		YearOfRelease: "2022",
// 	})
// 	movies = append(movies, &models.Movie{
// 		Id: primitive.ObjectID{},
// 		Name: "movie2",
// 		YearOfRelease: "2022",
// 	})
// 	movies = append(movies, &models.Movie{
// 		Id: primitive.ObjectID{},
// 		Name: "movie3",
// 		YearOfRelease: "2022",
// 	})
// 	movies = append(movies, &models.Movie{
// 		Id: primitive.ObjectID{},
// 		Name: "movie4",
// 		YearOfRelease: "2022",
// 	})
// 	movies = append(movies, &models.Movie{
// 		Id: primitive.ObjectID{},
// 		Name: "movie5",
// 		YearOfRelease: "2022",
// 	})

// 	t.Execute(w, movies)
// }
