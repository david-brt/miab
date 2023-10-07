package controllers

import (
	"database/sql"
	"messageinabottle/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	cost := 12
	hashedSaltedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	statement := `INSERT INTO User_ (username, password_hash_salted) VALUES ($1, $2)`

	_, err = db.Exec(statement, user.Username, string(hashedSaltedPassword))
	if err != nil {
		return c.Status(500).JSON("Could not insert into database.")
	}
	return c.JSON("User created successfully.")
}

// generates jwt token and returns it to client
func LoginHandler(c *fiber.Ctx, db *sql.DB) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var hashedSaltedPassword string
	statement := `SELECT password_hash_salted FROM USER_ WHERE username = $1`
	row := db.QueryRow(statement, user.Username)
	err := row.Scan(&hashedSaltedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedSaltedPassword), []byte(user.Password))

	if err != nil {
		return c.Status(401).JSON("Invalid password.")
	}

	expirationDate := time.Now().Add(time.Hour * 24 * 14)
	expirationDateNumeric := jwt.NumericDate{Time: expirationDate}
	currentTimeNumeric := jwt.NumericDate{Time: time.Now()}

	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  user.ID,
			"name": user.Username,
			"exp":  expirationDateNumeric,
			"iat":  currentTimeNumeric,
		})
	s, err := t.SignedString(key)
	if err != nil {
		return c.Status(500).JSON("An error occurred generating authentication token.")
	}

	return c.JSON(s)
}
