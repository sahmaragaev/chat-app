package db

import (
    "chat-app/pkg/models"
    "chat-app/pkg/utils"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// CreateMessage adds a new message to the database.
func CreateMessage(message *models.Message) (*mongo.InsertOneResult, error) {
    result, err := MessageCollection.InsertOne(context.Background(), message)
    if err != nil {
        utils.Log.Errorf("CreateMessage: Error inserting new message: %v", err)
        return nil, err
    }

    utils.Log.Infof("CreateMessage: Message created successfully: %v", result.InsertedID)
    return result, nil
}

// GetMessage retrieves a message by its ID.
func GetMessage(messageID primitive.ObjectID) (*models.Message, error) {
    var message models.Message
    if err := MessageCollection.FindOne(context.Background(), bson.M{"_id": messageID}).Decode(&message); err != nil {
        utils.Log.Errorf("GetMessage: Error finding message: %v", err)
        return nil, err
    }

    utils.Log.Infof("GetMessage: Message retrieved: %v", messageID)
    return &message, nil
}

func UpdateMessage(messageID primitive.ObjectID, updatedData bson.M) error {
    _, err := MessageCollection.UpdateOne(context.Background(), bson.M{"_id": messageID}, bson.M{"$set": updatedData})
    if err != nil {
        utils.Log.Errorf("UpdateMessage: Error updating message: %v", err)
        return err
    }

    utils.Log.Infof("UpdateMessage: Message updated successfully: %v", messageID)
    return nil
}

func DeleteMessage(messageID primitive.ObjectID) error {
    _, err := MessageCollection.DeleteOne(context.Background(), bson.M{"_id": messageID})
    if err != nil {
        utils.Log.Errorf("DeleteMessage: Error deleting message: %v", err)
        return err
    }

    utils.Log.Infof("DeleteMessage: Message deleted successfully: %v", messageID)
    return nil
}
