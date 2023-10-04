package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"messageinabottle/models/compounds"
)

func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {
	message := compounds.UserMessage{}

	if err := c.BodyParser(&message); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	userStatement := `INSERT INTO User_ (username)
		                VALUES ($1)`
	_, err := db.Exec(userStatement, message.SenderName)
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}

	messageStatement := `INSERT INTO Message (content, sender, timestamp)
		                VALUES ($1, (SELECT id FROM User_ WHERE username = $2), CURRENT_TIMESTAMP)`
	_, err = db.Exec(messageStatement, message.Content, message.SenderName)
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}

	return c.Status(200).JSON("Message sent successfully.")
}
