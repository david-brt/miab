package routes

import (
	"database/sql"
	"messageinabottle/handlers"

	"github.com/gofiber/fiber/v2"
)

func Authorized(db *sql.DB, router fiber.Router) {
	router.Get("/verify-token", func(c *fiber.Ctx) error {
		return handlers.VerifyTokenHandler(c)
	})

	router.Post("/rename", func(c *fiber.Ctx) error {
		return handlers.RenameHandler(c, db)
	})
}
