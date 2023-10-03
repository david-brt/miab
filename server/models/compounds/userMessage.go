package compounds

import (
	"time"
)

type UserMessage struct {
	MessageID  int       `json:"messageId"`
	Content    string    `json:"messageContent"`
	Timestamp  time.Time `json:"messageTimestamp"`
	SenderID   int       `json:"senderId"`
	SenderName string    `json:"senderName"`
}
