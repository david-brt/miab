package controllers

import (
	"database/sql"
	"log"
	"messageinabottle/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// generates jwt token and returns it to client
func LoginHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
    log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Wrong format",
    })
	}

	var hashedSaltedPassword string
	statement := `SELECT password_hash_salted FROM USER_ WHERE username = $1`
	row := db.QueryRow(statement, user.Username)
	err := row.Scan(&hashedSaltedPassword)

  if err != nil {
    log.Default().Println(err.Error())
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Could not process entity",
    })
  }

	err = bcrypt.CompareHashAndPassword([]byte(hashedSaltedPassword), []byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "info": "Invalid password",
    })
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
    log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Could not generate auth token",
    })
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
    "auth_token": s,
  })
}
