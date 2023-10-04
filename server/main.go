package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
	"messageinabottle/controllers"
	"messageinabottle/services"
	"os"
)

func main() {
	app := fiber.New()
	connStr := "postgresql://postgres:postgres@db/messageinabottle?sslmode=disable"
	db := services.ConnectDB(connStr)

	app.Get("/", func(c *fiber.Ctx) error {
		return controllers.IndexHandler(c, db)
	})
	app.Post("/update", func(c *fiber.Ctx) error {
		return controllers.UpdateHandler(c, db)
	})
	app.Post("/signup", func(c *fiber.Ctx) error {
		return controllers.SignupHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
