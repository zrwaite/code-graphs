package db

import (
	"context"
	"log"

	"github.com/zrwaite/github-graphs/db/validators"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDatabase() {
	MongoDatabase.Collection("users").Drop(context.TODO())
	var options = new(options.CreateCollectionOptions)
	options.Validator = validators.UsersValidator
	error := MongoDatabase.CreateCollection(context.TODO(), "users", options)
	if error != nil {
		log.Fatal("Failed to create users collection" + error.Error())
	}
}
