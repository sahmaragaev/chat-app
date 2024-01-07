package handlers

import (
	"chat-app/pkg/db"
	"chat-app/pkg/models"
	"chat-app/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		utils.Log.WithError(err).Error("Error decoding message data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.CreateMessage(&message)
	if err != nil {
		utils.Log.WithError(err).Error("Error creating new message")
		http.Error(w, "Error creating message", http.StatusInternalServerError)
		return
	}

	utils.Log.WithField("messageID", result.InsertedID).Info("Message created successfully")
	json.NewEncoder(w).Encode(result)
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.Log.WithError(err).Error("Error parsing message ID")
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	message, err := db.GetMessage(messageID)
	if err != nil {
		utils.Log.WithError(err).Error("Error retrieving message")
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	utils.Log.WithField("messageID", messageID).Info("Message retrieved successfully")
	json.NewEncoder(w).Encode(message)
}

func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.Log.WithError(err).Error("Error parsing message ID")
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var updatedData models.Message
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.Log.WithError(err).Error("Error decoding updated message data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	update := utils.MessageToBsonM(updatedData)
	err = db.UpdateMessage(messageID, update)
	if err != nil {
		utils.Log.WithError(err).Error("Error updating message")
		http.Error(w, "Error updating message", http.StatusInternalServerError)
		return
	}

	utils.Log.WithField("messageID", messageID).Info("Message updated successfully")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.Log.WithError(err).Error("Error parsing message ID")
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteMessage(messageID)
	if err != nil {
		utils.Log.WithError(err).Error("Error deleting message")
		http.Error(w, "Error deleting message", http.StatusInternalServerError)
		return
	}

	utils.Log.WithField("messageID", messageID).Info("Message deleted successfully")
	w.WriteHeader(http.StatusNoContent)
}
