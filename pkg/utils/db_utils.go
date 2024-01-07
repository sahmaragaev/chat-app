package utils

import (
	"chat-app/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func UserToBsonM(user models.User) (bson.M, error) {
	update := bson.M{}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		update["password"] = string(hashedPassword)
	}
	if user.DisplayName != "" {
		update["displayName"] = user.DisplayName
	}
	if user.Email != "" {
		update["email"] = user.Email
	}
	if !user.CreatedAt.IsZero() {
		update["createdAt"] = user.CreatedAt
	}
	update["active"] = user.Active

	return update, nil
}

func MessageToBsonM(message models.Message) bson.M {
	update := bson.M{}

	if message.SenderID != "" {
		update["senderId"] = message.SenderID
	}
	if message.Content != "" {
		update["content"] = message.Content
	}
	if !message.SentAt.IsZero() {
		update["sentAt"] = message.SentAt
	}
	if message.RoomID != "" {
		update["roomId"] = message.RoomID
	}

	return update
}
