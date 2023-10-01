package controllers

import (
	"database/sql"
	"log"
	"github.com/gofiber/fiber/v2"
  "messageinabottle/models"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
  message := models.Message{}
  row  := db.QueryRow("SELECT id, content, sender, timestamp FROM message ORDER BY timestamp DESC")

  err := row.Scan(&message.ID, &message.Content, &message.Sender, &message.Timestamp)
  if err != nil {
    log.Fatalln("No rows returned.")
  }

  return c.JSON(message)
}

func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {
  message := models.Message{}

  if err := c.BodyParser(&message); err != nil {
    return c.Status(400).JSON(err.Error())
  }

  statement := `INSERT INTO message (content, sender, timestamp) 
                VALUES ($1, $2, CURRENT_TIMESTAMP)`

  _, err := db.Exec(statement, message.Content, message.Sender)

  if err != nil {
    return c.Status(400).JSON("Could not insert into database.")
  }

  return c.Status(200).JSON("Message sent successfully.")
}
