package handlers

import (
	"database/sql"
	"messageinabottle/dataaccess"
	"messageinabottle/errors"
	"messageinabottle/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// generates jwt token and returns it to client
func LoginHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return errors.ParsingError(c, err)
	}

	if dataaccess.UserExists(db, user.Username) && dataaccess.IsOnCooldown(db, &user) {
		return errors.TooManyRequestsError(c)
	}

  userID,err := dataaccess.GetUserId(db, user.Username)
	user.ID = userID

	if !dataaccess.PasswordMatchesHash(db, &user) {
		dataaccess.UpdateAttemptStatus(db, &user, true)
		return errors.UnauthorizedError(c)
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
		return errors.InternalServerError(c, err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    s,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
	})

	dataaccess.UpdateAttemptStatus(db, &user, false)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": "true",
		"user": fiber.Map{
			"id":   user.ID,
			"name": user.Username,
		},
	})
}
