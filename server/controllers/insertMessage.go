package controllers

import (
	"database/sql"
	"messageinabottle/models"

	"github.com/gofiber/fiber/v2"
)

func InsertMessageHandler(c *fiber.Ctx, db *sql.DB) error {
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	messageStatement := `INSERT INTO Message (content, sender, sender_name, timestamp)
		                VALUES ($1, $2, $3, CURRENT_TIMESTAMP)`

	var err error
	if message.ID == 0 {
		_, err = db.Exec(messageStatement, message.Content, sql.NullString{}, message.SenderName)
	} else {
		_, err = db.Exec(messageStatement, message.Content, message.ID, message.SenderName)
	}

	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}

	return c.Status(200).JSON("Message sent successfully.")
}