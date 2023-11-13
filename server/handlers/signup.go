package handlers

import (
	"database/sql"
	"log"
	"messageinabottle/dataaccess"
	"messageinabottle/models"
	"messageinabottle/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		log.Default().Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Rquest Body",
      "errorMessage": "Something went wrong, try again later",
		})
	}

	if !utils.IsValidPassword(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid password",
      "errorMessage": "Password must contain at least eight characters including an uppercase and a lowercase letter, a number and a special character (#?!@$ %^&*)",
		})
	}

	if !utils.IsValidUsername(user.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid username",
			"errorMessage": "Please only use letters, numbers and underscores",
		})
	}

	if dataaccess.UserExists(db, user.Username) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User exists",
			"errorMessage": "Username is already taken",
		})
	}

	cost := 12
	hashedSaltedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not process entity",
			"errorMessage": "Something went wrong, try again later",
		})
	}

	statement := `INSERT INTO User_ (username, password_hash_salted, latest_login_attempt, failed_login_attempts) VALUES ($1, $2, CURRENT_TIMESTAMP, 0)`

	_, err = db.Exec(statement, user.Username, string(hashedSaltedPassword))
	if err != nil {
		log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not process entity",
			"errorMessage": "Something went wrong, try again later",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "User created",
	})
}
