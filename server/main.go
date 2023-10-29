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

	verficationMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(os.Getenv("JWT_SECRET")),
			JWTAlg: jwtware.HS256,
		},
	})
	cors := cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	})

	app := fiber.New()
	api := app.Group("/api", cors, logger.New())
	authorized := api.Group("/authorized", verficationMiddleware)

	connStr := "postgresql://postgres:postgres@db/messageinabottle?sslmode=disable"
	db := services.ConnectDB(connStr)

	log.SetFlags(log.Lshortfile)

	api.Get("/global-message", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(c, db)
	})

	api.Get("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie("auth_token")
		return nil
	})

	api.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, db)
	})

	api.Post("/send-message", func(c *fiber.Ctx) error {
		return handlers.InsertMessageHandler(c, db)
	})

	authorized.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignupHandler(c, db)
	})

	authorized.Get("/verify-token", func(c *fiber.Ctx) error {
		return handlers.VerifyTokenHandler(c)
	})

	authorized.Post("/rename", func(c *fiber.Ctx) error {
		return handlers.RenameHandler(c, db)
	})

	port := os.Getenv("PORT")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
