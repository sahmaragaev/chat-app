package db

import (
	"chat-app/pkg/models"
	"chat-app/pkg/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Log.Errorf("CreateUser: Error hashing password: %v", err)
		return nil, err
	}
	user.Password = hashedPassword

	result, err := UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		utils.Log.Errorf("CreateUser: Error inserting new user: %v", err)
		return nil, err
	}

	utils.Log.Infof("CreateUser: User created successfully: %v", result.InsertedID)
	return result, nil
}

func GetUserById(userID primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := UserCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user); err != nil {
		utils.Log.Errorf("GetUser: Error finding user: %v", err)
		return nil, err
	}

	utils.Log.Infof("GetUser: User retrieved: %v", userID)
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		utils.Log.Errorf("GetUser: Error finding user by email: %v", err)
		return nil, err
	}
	utils.Log.Infof("GetUser: User retrieved: %v", email)
	return &user, nil
}

func UpdateUser(userID primitive.ObjectID, updateData bson.M) error {
	if updateData["password"] != nil {
		hashedPassword, err := utils.HashPassword(updateData["password"].(string))
		if err != nil {
			utils.Log.Errorf("UpdateUser: Error hashing password: %v", err)
			return err
		}
		updateData["password"] = hashedPassword
	}

	_, err := UserCollection.UpdateOne(context.Background(), bson.M{"_id": userID}, bson.M{"$set": updateData})
	if err != nil {
		utils.Log.Errorf("UpdateUser: Error updating user: %v", err)
		return err
	}

	utils.Log.Infof("UpdateUser: User updated successfully: %v", userID)
	return nil
}

func DeleteUser(userID primitive.ObjectID) error {
	_, err := UserCollection.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		utils.Log.Errorf("DeleteUser: Error deleting user: %v", err)
		return err
	}

	utils.Log.Infof("DeleteUser: User deleted successfully: %v", userID)
	return nil
}
