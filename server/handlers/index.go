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

	// check if there is a message that has not been read yet
	row := db.QueryRow("SELECT id, content, sender, sender_name, timestamp FROM Message Where has_been_read IS FALSE ORDER BY timestamp ASC")
	err := row.Scan(&message.ID, &message.Content, &sender, &message.SenderName, &message.Timestamp)


	// if every message was already read by a user, send the newest message
	if err != nil {
		row = db.QueryRow("SELECT id, content, sender, sender_name, timestamp FROM Message ORDER BY timestamp DESC")
		err = row.Scan(&message.ID, &message.Content, &sender, &message.SenderName, &message.Timestamp)
	}

	// if here are no messages... idk, this
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

	// mark the message that is beeing sent to the user as has_been_read ✔️

	sqlStatement := "Update message SET has_been_read = true where id = $1;"

	_, err2 := db.Exec(sqlStatement, message.ID)

	if err2 != nil {
		log.Default().Println(err2.Error())
    }


	return c.JSON(message)
}
