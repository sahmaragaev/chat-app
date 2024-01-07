package routes

import (
	"chat-app/pkg/handlers"
	"chat-app/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
    userRouter := router.PathPrefix("/users").Subrouter()

    userRouter.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    userRouter.HandleFunc("/login", handlers.LoginUser).Methods("POST")

    protectedUserRouter := userRouter.NewRoute().Subrouter()
    protectedUserRouter.Use(middlewares.JWTAuthentication)
    protectedUserRouter.HandleFunc("/{id}", handlers.GetUserHandler).Methods("GET")
    protectedUserRouter.HandleFunc("/{id}", handlers.UpdateUserHandler).Methods("PUT")
    protectedUserRouter.HandleFunc("/{id}", handlers.DeleteUserHandler).Methods("DELETE")
}
