package db

import (
	"chat-app/pkg/utils"
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(connectionString string) error {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		utils.Log.Fatalf("Error connecting to MongoDB: %v", err)
		return err
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		utils.Log.Fatalf("Error pinging MongoDB: %v", err)
		return err
	}

	utils.Log.Info("Connected to MongoDB!")
	return nil
}