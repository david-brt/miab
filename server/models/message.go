package models

import (
	"time"
)

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"message"`
	SenderID   int       `json:"senderId"`
	SenderName string    `json:"sender"`
	Timestamp  time.Time `json:"timestamp"`
}
