package handlers

import (
  	"github.com/gofiber/fiber/v2"
    "time"
 )

func LogoutHandler(c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
    Name: "auth_token",
    Expires: time.Now().Add(-(2 * time.Hour)),
    HTTPOnly: true,
  })
  return nil
}
