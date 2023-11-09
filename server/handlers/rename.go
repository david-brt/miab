package handlers

import (
	"database/sql"
	"log"
	"messageinabottle/utils"

	"github.com/gofiber/fiber/v2"
)

func RenameHandler(c *fiber.Ctx, db *sql.DB) error {

	var requestData struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		log.Printf("Fehler beim Parsen des Request-Bodys: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ungültiger Request-Body",
		})
	}

	if !utils.IsValidUsername(requestData.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid username",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "User renamed",
	  })

}
