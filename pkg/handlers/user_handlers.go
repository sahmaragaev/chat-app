package handlers

import (
	"chat-app/pkg/db"
	"chat-app/pkg/models"
	"chat-app/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Log.WithError(err).Error("Error decoding user data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Log.WithError(err).Error("Error hashing password")
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	_, err = db.CreateUser(&user)
	if err != nil {
		utils.Log.WithError(err).Error("Error creating user")
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{"userID": userIDStr}).WithError(err).Error("Error converting userID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserById(userID)
	if err != nil {
		utils.Log.WithError(err).Error("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("id")
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        utils.Log.WithError(err).Error("Error converting userID")
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.Log.WithError(err).Error("Error decoding update data")
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    update, err := utils.UserToBsonM(user)
    if err != nil {
        utils.Log.WithError(err).Error("Error converting user to bson.M")
        http.Error(w, "Error processing user data", http.StatusInternalServerError)
        return
    }

    err = db.UpdateUser(userID, bson.M{"$set": update})
    if err != nil {
        utils.Log.WithError(err).Error("Error updating user")
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{"userID": userIDStr}).WithError(err).Error("Error converting userID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteUser(userID)
	if err != nil {
		utils.Log.WithError(err).Error("Error deleting user")
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}