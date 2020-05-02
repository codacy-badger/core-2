package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// Connect DB instantiates a mongoDB connection
func ConnectDB() {
	var mongoURL = os.Getenv("IBORG_MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(mongoURL)
	var err error
	DBConn, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Connected to MongoDB")
}
