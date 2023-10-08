package controllers

import (
	"database/sql"
	"log"
	"messageinabottle/models"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	message := models.Message{}

	var sender sql.NullInt32
	row := db.QueryRow("SELECT id, content, sender, sender_name, timestamp FROM Message ORDER BY timestamp DESC")
	err := row.Scan(&message.ID, &message.Content, &sender, &message.SenderName, &message.Timestamp)

	if err != nil {
		log.Fatalln("No rows returned.")
		return c.Status(500).JSON("error fetching from database")
	}

	if sender.Valid {
		message.Sender = int(sender.Int32)
	} else {
		message.Sender = -1
	}

	return c.JSON(message)
}
