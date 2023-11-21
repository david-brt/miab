package errors

import "github.com/gofiber/fiber/v2"

func UnauthorizedError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":        "Invalid credentials",
		"errorMessage": "Invalid credentials",
	})
}
