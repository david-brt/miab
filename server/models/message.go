package models

import (
	"time"
)

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Sender     int       `json:"sender"`
	SenderName string    `json:"senderName"`
	Timestamp  time.Time `json:"timestamp"`
}
