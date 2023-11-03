package routes

import (
	"database/sql"
	"messageinabottle/handlers"

	"github.com/gofiber/fiber/v2"
)

func Unauthorized(db *sql.DB, router fiber.Router) {
	router.Get("/global-message", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(c, db)
	})

	router.Get("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie("auth_token")
		return nil
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, db)
	})

	router.Post("/send-message", func(c *fiber.Ctx) error {
		return handlers.InsertMessageHandler(c, db)
	})

	router.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignupHandler(c, db)
	})

}
