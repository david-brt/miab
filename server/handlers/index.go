package handlers

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
		log.Default().Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "No message found",
    })
	}

  // valid if user is logged in
	if sender.Valid {
		message.Sender = int(sender.Int32)
	} else {
		message.Sender = -1
	}

	return c.JSON(message)
}
