package db_service

import (
	"context"

	"github.com/zrwaite/github-graphs/db"
	"github.com/zrwaite/github-graphs/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers() ([]*models.User, error) {
	cursor, err := db.MongoDatabase.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUser(user *models.User) error {
	_, err := db.MongoDatabase.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}