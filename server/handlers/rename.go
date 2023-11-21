package handlers

import (
	"database/sql"
	"messageinabottle/errors"
	"messageinabottle/utils"

	"github.com/gofiber/fiber/v2"
)

func RenameHandler(c *fiber.Ctx, db *sql.DB) error {

	var requestData struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return errors.ParsingError(c, err)
	}
	log.Default().Println(requestData)

	if !utils.IsValidUsername(requestData.Username) {
		return errors.MalformedUsernameError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "User renamed",
	})

}
