package errors

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func NotFoundError(c *fiber.Ctx, err error) error {
	log.Default().Println(err.Error())
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":        "Resource not found",
		"errorMessage": "Could not find resource",
	})
}
