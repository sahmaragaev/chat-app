package main

import (
	"chat-app/pkg/db"
	"chat-app/pkg/routes"
	"chat-app/pkg/utils"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.Log.Fatalf("Error loading .env file: %v", err)
	}

	connectionString := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	if err := db.Connect(connectionString); err != nil {
		utils.Log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	db.InitCollections(db.Client, dbName)

	router := mux.NewRouter()

	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterMessageRoutes(router)
}
