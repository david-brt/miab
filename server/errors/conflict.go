package errors

import (
	"github.com/gofiber/fiber/v2"
)

// Returns HTTP status 409 Conflict.
//
// errorMessage: "Username is already taken"
func UserExistsError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"error":        "User exists",
		"errorMessage": "User already exists",
	})
}
