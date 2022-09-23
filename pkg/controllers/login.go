package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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
	if err != nil {
		log.Fatal(err)
	}

	// store the user data in mongodb inside the user collection

	mongoDBClient := connections.MongoDBClient

	user := &models.User{
		Name:     signUpData.Name,
		Email:    signUpData.Email,
		Password: hashedUserPassword,
	}

	// connect to db gotofasal
	gotofasalDBDatabase := mongoDBClient.Database("gotofasal")
	userCollection := gotofasalDBDatabase.Collection("user")
	
	// check if user already registered.
	findOneCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	check := userCollection.FindOne(findOneCtx, &models.User{Email:user.Email})
	fmt.Println(check)


	insertOneCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	dbInsertOneResult, err := userCollection.InsertOne(insertOneCtx, user)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Inserted the following documents in user Collection of gotofasal db :\n%v\n", dbInsertOneResult.InsertedID)

	http.Redirect(w, request, "/signin", http.StatusSeeOther)
	
	return
}

func SignIn(w http.ResponseWriter, request *http.Request) {
	return
}
