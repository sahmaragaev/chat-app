package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID       primitive.ObjectID `json:"id"`
	SenderID string             `json:"senderId"`
	Content  string             `json:"content"`
	SentAt   time.Time          `json:"sentAt"`
	RoomID   string             `json:"roomId"`
}
