package errors

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// Returns HTTP status 500 InternalServerError.
//
// errorMessage: "Something went wrong, try again later"
func InternalServerError(c *fiber.Ctx, err error) error {
	log.Default().Println(err.Error())
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":        "Could not process entity",
		"errorMessage": "Something went wrong, try again later",
	})
}
