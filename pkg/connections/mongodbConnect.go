package connections

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoDBClient *mongo.Client
var MongoDBCtx *context.Context

func ConnectToMongoDB() {

	var ctx = context.TODO()
	mongoDBClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoDBClient, err := mongo.Connect(ctx, mongoDBClientOptions)
	fmt.Println("Connecting to mongodb")
	if err != nil {
		log.Fatal(err)
	}
	err = mongoDBClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	databaseNames, err := mongoDBClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Databases: %v\n", databaseNames)

	MongoDBClient = mongoDBClient
	MongoDBCtx = &ctx
}
