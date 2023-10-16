package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "info": "User not logged in",
    })
	}

	return c.JSON(fiber.Map{
    "claims": user.Claims,
  })
}
