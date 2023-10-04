package controllers

import (
	"database/sql"
	"log"
	"messageinabottle/models/compounds"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	message := compounds.UserMessage{}

	row := db.QueryRow("SELECT m.id, m.content, m.timestamp, u.id, u.username FROM Message m JOIN User_ u ON m.sender = u.id ORDER BY timestamp DESC")
	err := row.Scan(&message.MessageID, &message.Content, &message.Timestamp, &message.SenderID, &message.SenderName)

	if err != nil {
		log.Fatalln("No rows returned.")
		return c.Status(500).JSON("error fetching from database")
	}

	return c.JSON(message)
}
