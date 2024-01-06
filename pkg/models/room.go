package models

import "time"

type Room struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    Members     []string  `json:"members"`
}
