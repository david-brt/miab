package models

import (
  "time"
)

type Message struct {
  ID        string     `json:"id"`
  Content   string     `json:"content"`
  Sender    string     `json:"sender"`
  Timestamp time.Time  `json:"timestamp"`
}

