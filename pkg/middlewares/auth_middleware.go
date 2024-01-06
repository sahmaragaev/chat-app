package middleware

import (
	"chat-app/pkg/utils"
	"context"
	"log"
	"net/http"
	"strings"
)

func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("JWTAuthentication: Missing Authorization header")
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			log.Println("JWTAuthentication: Invalid Authorization header format")
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]
		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Printf("JWTAuthentication: Invalid token: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
