package handlers

import (
	"database/sql"
	"log"
	"messageinabottle/dataaccess"
	"messageinabottle/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RenameHandler(c *fiber.Ctx, db *sql.DB) error {
	
}