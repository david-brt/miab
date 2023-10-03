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

func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {
	message := compounds.UserMessage{}

	if err := c.BodyParser(&message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userStatement := `INSERT INTO User_ (username)
		                VALUES ($1)`
	_, err := db.Exec(userStatement, message.SenderName)
	if err != nil {
		return c.Status(400).JSON("Could not insert into database.")
	}

	messageStatement := `INSERT INTO Message (content, sender, timestamp)
		                VALUES ($1, (SELECT id FROM User_ WHERE username = $2), CURRENT_TIMESTAMP)`
	_, err = db.Exec(messageStatement, message.Content, message.SenderName)
	if err != nil {
		return c.Status(400).JSON("Could not insert into database.")
	}

	return c.Status(200).JSON("Message sent successfully.")
}
