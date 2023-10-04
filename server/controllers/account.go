package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"messageinabottle/models"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	statement := `INSERT INTO User_ (username) VALUES ($1)`

	_, err := db.Exec(statement, user.Username)
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}
	return c.Status(200).JSON("User created successfully.")
}
