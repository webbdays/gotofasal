package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/webbdays/gotofasal/pkg/connections"
	"github.com/webbdays/gotofasal/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, request *http.Request) {
	var signUpData *models.SignUpData

	// read/parse the request body from the http request
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to struct from json and validate the request body
	_ = json.Unmarshal(requestBody, &signUpData)
	if err := validator.New().Struct(signUpData); err != nil {
		log.Fatal(err)
	}

	// hash the password
	hashedUserPassword, err := bcrypt.GenerateFromPassword([]byte(signUpData.Password), 15)
	if err !=nil {
		log.Fatal(err)
	}

	// store the user data in mongodb inside the user collection
	mongoDBClient := connections.MongoDBClient
	mongoDBCtx := connections.MongoDBCtx


	user := &models.User{
		Name:     signUpData.Name,
		Email:    signUpData.Email,
		Password: hashedUserPassword,
	}

	gotofasalDBDatabase := mongoDBClient.Database("gotofasal")
	userCollection := gotofasalDBDatabase.Collection("user")
	dbInsertOneResult, err := userCollection.InsertOne(*mongoDBCtx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted the following documents in user Collection of gotofasal db :\n%v\n", dbInsertOneResult.InsertedID)
}




func SignIn(w http.ResponseWriter, request *http.Request) {
	return
}
