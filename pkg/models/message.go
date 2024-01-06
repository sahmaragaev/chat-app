package models

import "time"

type Message struct {
    ID        string    `json:"id"`
    SenderID  string    `json:"senderId"`
    Content   string    `json:"content"`
    SentAt    time.Time `json:"sentAt"`
    RoomID    string    `json:"roomId"`
}
