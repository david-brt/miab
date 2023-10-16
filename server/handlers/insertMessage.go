package handlers

import (
	"database/sql"
	"log"
	"messageinabottle/models"

	"github.com/gofiber/fiber/v2"
)

func InsertMessageHandler(c *fiber.Ctx, db *sql.DB) error {
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil {
    log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Wrong format",
    })
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
    log.Default().Println(err.Error())
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Could not insert message",
    })
  }

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
    "success": "Message sent",
  })
}
