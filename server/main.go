package main

import (
	"fmt"
	"log"
	"messageinabottle/handlers"
	"messageinabottle/services"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
  app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:5173",
    AllowHeaders:  "Origin, Content-Type, Accept",
    AllowCredentials: true,
  }))

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(c, db)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, db)
	})

	app.Post("/send-message", func(c *fiber.Ctx) error {
		return handlers.InsertMessageHandler(c, db)
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignupHandler(c, db)
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(os.Getenv("JWT_SECRET")),
			JWTAlg: jwtware.HS256,
		},
	}))

	app.Get("/verify-token", func(c *fiber.Ctx) error {
		return handlers.VerifyTokenHandler(c)
	})

	port := os.Getenv("PORT")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
