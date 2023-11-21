package dataaccess

import (
	"database/sql"
	"log"
	"messageinabottle/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// takes a sql.DB object and a username as a string
func UserExists(db *sql.DB, username string) bool {
	var exists bool
	row := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM User_ WHERE username = $1)`, username)
	err := row.Scan(&exists)

	if err != nil {
		return false
	}
	return exists
}

// - returns a bool indicating whether a user is on cooldown
// - cooldown (10s) after 5 unsuccessful attempts
func IsOnCooldown(db *sql.DB, user *models.User) bool  {
  var lastAttempt time.Time
  var failedAttempts int

  statement := `SELECT latest_login_attempt, failed_login_attempts FROM USER_ WHERE username = $1`
  row := db.QueryRow(statement, user.Username)
  err := row.Scan(&lastAttempt, &failedAttempts)
  if err != nil {
    log.Default().Println(err.Error())
  }

  secondsSinceLastAttempt := time.Now().Sub(lastAttempt).Seconds()

  if failedAttempts > 5 && secondsSinceLastAttempt < 10 {
    return true
  }
  return false
}

// - updates latest_login_attempt to CURRENT_TIMESTAMP and increments failed_login_attempts by 1
// - if increment is true, adds 1 to failed_login_attempts, else resets to 0
func UpdateAttemptStatus(db *sql.DB, user *models.User, increment bool) {
  statement := `UPDATE User_ SET latest_login_attempt = CURRENT_TIMESTAMP, failed_login_attempts = 0 WHERE username = $1`
  if increment {
    statement = `UPDATE User_ SET latest_login_attempt = CURRENT_TIMESTAMP, failed_login_attempts = failed_login_attempts + 1 WHERE username = $1`
  }

  _, err := db.Exec(statement, user.Username)
  if err != nil {
    log.Default().Println(err.Error())
    // further error handling might be required
  }
}

// checks if salted password matches hash stored in db
func PasswordMatchesHash(db *sql.DB, user *models.User) bool {
	var hashedSaltedPassword string
  statement := `SELECT password_hash_salted FROM USER_ WHERE username = $1`
	row := db.QueryRow(statement, user.Username)
	err := row.Scan(&hashedSaltedPassword)
  if err != nil {
    log.Default().Println()
    return false
  }
  err = bcrypt.CompareHashAndPassword([]byte(hashedSaltedPassword), []byte(user.Password))
  if err != nil {
    log.Default().Println()
    return false
  }
  return true
}

func GetUserId(db *sql.DB, username string) (int, error){
  statement := `Select id from user_ where username = $1; `
  
  // Execute QueryRow to retrieve a single row result
  row := db.QueryRow(statement, username)

  // Scan the retrieved ID into the variable
	var userID int
    err := row.Scan(&userID)
    if err != nil {
        return 0,err
    }
  return userID,nil
}