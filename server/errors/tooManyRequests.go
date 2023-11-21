package errors

import "github.com/gofiber/fiber/v2"

// Returns HTTP status 429 TooManyRequests.
//
// errorMessage: "Please wait a few moments before trying again"
func TooManyRequestsError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"error":        "On cooldown",
		"errorMessage": "Please wait a few moments before trying again",
	})
}
