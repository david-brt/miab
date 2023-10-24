package handlers

import (
	"database/sql"
  "errors"
	"log"
	"messageinabottle/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// - returns a bool indicating whether a user is on cooldown and a float that holds the seconds until not on cooldown anymore
// - cooldown (10s) gets activated after 5 unsuccessful attempts
func isOnCooldown(user *models.User, db *sql.DB) (bool, float) {
  var lastAttempt time.Time
  var failedAttempts int

  statement := `SELECT latest_login_attempt, failed_login_attempts FROM USER_ WHERE username = $1`
  row := db.QueryRow(statement, user.Username)
  err := row.Scan(&lastAttempt, &failedAttempts)
  if err != nil {
    log.Default().Println(err.Error())
  }

  secondsSinceLastAttempt := time.Now().Sub(lastAttempt).Seconds()

  if failedAttempts <= 5 {
    return false, 0
  }
  if secondsSinceLastAttempt <= 10 {
    return true, 10 - secondsSinceLastAttempt
  }
  return false, 0
}

// generates jwt token and returns it to client
func LoginHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
    log.Default().Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Wrong format",
    })
	}

  isOnCooldown, onCooldownFor := isOnCooldown(user, db)

  if isOnCooldown {
    return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
      "error": "On cooldown",
      "onCooldownFor": onCooldownFor,
    })
  }

	var hashedSaltedPassword string
	statement := `SELECT password_hash_salted FROM USER_ WHERE username = $1`
	row := db.QueryRow(statement, user.Username)
	err := row.Scan(&hashedSaltedPassword)

  if err != nil {
    log.Default().Println(err.Error())
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "error": "Invalid credentials",
    })
  }

	err = bcrypt.CompareHashAndPassword([]byte(hashedSaltedPassword), []byte(user.Password))

	if err != nil {
    log.Default().Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "error": "Invalid credentials",
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

  c.Cookie(&fiber.Cookie{
    Name: "auth_token",
    Value: s,
    Expires: time.Now().Add(7 * 24 * time.Hour),
    HTTPOnly: true,
  })

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
    "success": "true",
    "user": fiber.Map{
      "id": user.ID,
      "name": user.Username,
    },
  })
