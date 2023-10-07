package controllers

import (
	"database/sql"
	"messageinabottle/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	statement := `INSERT INTO User_ (username) VALUES ($1)`

	_, err := db.Exec(statement, user.Username)
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}
	return c.JSON("User created successfully.")
}

// generates jwt token and returns it to client
func LoginHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	expirationDate := time.Now().Add(time.Hour * 24 * 14)
	expirationDateNumeric := jwt.NumericDate{Time: expirationDate}
	currentTimeNumeric := jwt.NumericDate{Time: time.Now()}

	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  user.ID,
			"name": user.Username,
			"exp":  expirationDateNumeric,
			"iat":  currentTimeNumeric,
		})
	s, err := t.SignedString(key)
	if err != nil {
		return c.Status(500).JSON("An error occurred generating authentication token.")
	}

	return c.JSON(s)
}
