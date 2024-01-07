package routes

import (
    "chat-app/pkg/handlers"
    "chat-app/pkg/middlewares"
    "github.com/gorilla/mux"
)

func RegisterMessageRoutes(router *mux.Router) {
    messageRouter := router.PathPrefix("/messages").Subrouter()

    messageRouter.Use(middlewares.JWTAuthentication)

    messageRouter.HandleFunc("/", handlers.CreateMessageHandler).Methods("POST")
    messageRouter.HandleFunc("/{id}", handlers.GetMessageHandler).Methods("GET")
    messageRouter.HandleFunc("/{id}", handlers.UpdateMessageHandler).Methods("PUT")
    messageRouter.HandleFunc("/{id}", handlers.DeleteMessageHandler).Methods("DELETE")
}
