package errors

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Intended for errors parsing the request body. Logs error and returns HTTP status 400 Bad Request.
//
// errorMessage: "Something went wrong, try again later"
func ParsingError(c *fiber.Ctx, err error) error {
	log.Default().Println(err.Error())
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":        "Invalid Request Body",
		"errorMessage": "Something went wrong, try again later",
	})
}

// Returns HTTP status 400 Bad Request.
//
// errorMessage: "Password must contain 8 to 50 characters including an uppercase and a lowercase letter, a number and a special character (#?!@$ %^&*)"
func MalformedPasswordError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":        "Malformed password",
		"errorMessage": "Password must contain 8 to 50 characters including an uppercase and a lowercase letter, a number and a special character (#?!@$ %^&*)",
	})
}

// Returns HTTP status 400 Bad Request.
//
// errorMessage: "Username must only contain letters, numbers and underscores"
func MalformedUsernameError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":        "Invalid username",
		"errorMessage": "Username must only contain letters, numbers and underscores",
	})
}
