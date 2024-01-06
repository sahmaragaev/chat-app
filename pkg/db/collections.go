package db

import "go.mongodb.org/mongo-driver/mongo"

var UserCollection *mongo.Collection
var MessageCollection *mongo.Collection

func InitCollections(client *mongo.Client, dbName string) {
	UserCollection = client.Database(dbName).Collection("users")
	MessageCollection = client.Database(dbName).Collection("messages")
}