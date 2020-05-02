package database

import (
	"context"
	"github.com/pravinba9495/iborg-core/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// SaveUser saves a user in the database
func SaveUser(user models.User) int {
	client := DBConn
	collection := client.Database("iborg").Collection("users")
	if collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}}).Err() != nil {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
			return 500
		}
		return 201
	}
	log.Println("Email already present")
	return 409
}
