package routes

import (
	"chat-app/pkg/db"
	"chat-app/pkg/handlers"
	"chat-app/pkg/middlewares"
	"chat-app/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterAuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middlewares.JWTAuthentication)
	authRouter.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	authRouter.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	authRouter.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	if userIDStr == "" {
		utils.Log.Error("User ID not found in request context")
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"userID": userIDStr,
			"error":  err,
		}).Error("Error converting userID to ObjectID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserById(userID)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"userID": userID,
			"error":  err,
		}).Error("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := fmt.Sprintf("Access granted. User ID: %s, User Email: %s", user.ID.Hex(), user.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
