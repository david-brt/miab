package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User is not authenticated",
		})
	}

	return c.JSON(user.Claims.(jwt.MapClaims))
}
