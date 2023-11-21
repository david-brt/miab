package handlers

import (
	"database/sql"
	"messageinabottle/dataaccess"
	"messageinabottle/errors"
	"messageinabottle/models"
	"messageinabottle/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return errors.ParsingError(c, err)
	}

	if !utils.IsValidPassword(user.Password) {
		return errors.MalformedPasswordError(c)
	}

	if !utils.IsValidUsername(user.Username) {
		return errors.MalformedUsernameError(c)
	}

	if dataaccess.UserExists(db, user.Username) {
		return errors.UserExistsError(c)
	}

	cost := 12
	hashedSaltedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return errors.InternalServerError(c, err)
	}

	statement := `INSERT INTO User_ (username, password_hash_salted, latest_login_attempt, failed_login_attempts) VALUES ($1, $2, CURRENT_TIMESTAMP, 0)`

	_, err = db.Exec(statement, user.Username, string(hashedSaltedPassword))
	if err != nil {
		return errors.InternalServerError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "User created",
	})
}
