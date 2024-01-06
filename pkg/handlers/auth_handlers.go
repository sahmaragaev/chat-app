package handlers

import (
    "chat-app/pkg/db"
    "chat-app/pkg/models"
    "chat-app/pkg/utils"
    "encoding/json"
    "net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var creds models.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := db.GetUserByEmail(creds.Email)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    err = utils.ComparePasswords(user.Password, creds.Password)
    if err != nil {
        http.Error(w, "Invwellalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.ID.Hex())
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    _, err = db.CreateUser(&user)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
