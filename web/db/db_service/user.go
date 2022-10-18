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

func GetUsersWithCache() ([]*models.User, error) {
	cacheKey := "users"
	users := []*models.User{}
	found := db.GetJsonCache(cacheKey, &users)
	if found {
		return users, nil
	}
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}
	db.SetJsonCache(cacheKey, users)
	return users, nil
}

func GetUser(username string) (*models.User, error) {
	cursor := db.MongoDatabase.Collection("users").FindOne(context.TODO(), bson.M{
		"username": username,
	})
	var user models.User
	if err := cursor.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(user *models.User) error {
	_, err := db.MongoDatabase.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	_, err := db.MongoDatabase.Collection("users").UpdateOne(context.TODO(), bson.M{
		"username": user.Username,
	}, bson.M{
		"$set": bson.M{
			"access_token": user.AccessToken,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
