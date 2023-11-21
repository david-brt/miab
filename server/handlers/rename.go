package handlers

import (
	"database/sql"
	"log"


	"github.com/gofiber/fiber/v2"

)

func RenameHandler(c *fiber.Ctx, db *sql.DB) error {

	var requestData struct {
		Username string `json:"username"`
	}

	// Überprüfe, ob der Request-Body ein JSON ist und parsen ihn
	if err := c.BodyParser(&requestData); err != nil {
		log.Printf("Fehler beim Parsen des Request-Bodys: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ungültiger Request-Body",
		})
	}
	log.Default().Println(requestData)

	// Du kannst jetzt auf den Nutzernamen zugreifen, der im Request gesendet wurde

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "User renamed",
	  })

}