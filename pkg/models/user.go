package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id"`
	DisplayName string             `json:"displayName"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	CreatedAt   time.Time          `json:"createdAt"`
	Active      bool               `json:"active"`
}
