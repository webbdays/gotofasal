package connections

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoDBClient *mongo.Client


func ConnectToMongoDB() {

	// var ctx = context.TODO()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	mongoDBClientOptions := options.Client().ApplyURI("mongodb+srv://Cluster0:" + os.Getenv("MONGODB_PASSWORD") + "@cluster0.gxamoya.mongodb.net/?retryWrites=true&w=majority")
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

	// databaseNames, err := mongoDBClient.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Databases: %v\n", databaseNames)

	MongoDBClient = mongoDBClient
}
