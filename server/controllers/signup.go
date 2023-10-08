package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"messageinabottle/dataaccess"
	"messageinabottle/models"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if dataaccess.UserExists(db, user.Username) {
		return c.Status(502).JSON("User already exists.")
	}

	cost := 12
	hashedSaltedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	statement := `INSERT INTO User_ (username, password_hash_salted) VALUES ($1, $2)`

	_, err = db.Exec(statement, user.Username, string(hashedSaltedPassword))
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}
	return c.JSON("User created successfully.")
}
