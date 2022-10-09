package validators

import "go.mongodb.org/mongo-driver/bson"

var bsonBoolean = bson.M{"bsonType": "bool"}
var bsonString = bson.M{"bsonType": "string"}

var UsersValidator = bson.M{
	"$jsonSchema": userSchema,
}

var userSchema = bson.M{
	"bsonType": "object",
	"required": []string{"username", "access_token", "refresh_token", "verified"},
	"properties": bson.M{
		"username":      bsonString,
		"access_token":  bsonString,
		"refresh_token": bsonString,
		"verified":      bsonBoolean,
	},
}
