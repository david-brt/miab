package main

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
  "github.com/gofiber/fiber/v2/middleware/logger"
	"messageinabottle/controllers"
	"messageinabottle/services"
	"os"
)

func main() {
	env := os.Getenv("MIAB_ENV")
	if env == "" {
		env = "development"
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env file.")
		}
	}

	app := fiber.New()
	connStr := "postgresql://postgres:postgres@db/messageinabottle?sslmode=disable"
	db := services.ConnectDB(connStr)

  app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return controllers.IndexHandler(c, db)
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return controllers.LoginHandler(c, db)
	})

	app.Post("/insert-message", func(c *fiber.Ctx) error {
		return controllers.InsertMessageHandler(c, db)
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		return controllers.SignupHandler(c, db)
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(os.Getenv("JWT_SECRET")),
			JWTAlg: jwtware.HS256,
		},
	}))

	app.Get("/verify-token", func(c *fiber.Ctx) error {
		return controllers.VerifyTokenHandler(c)
	})

	port := os.Getenv("PORT")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
