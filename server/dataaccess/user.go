package dataaccess

import (
	"database/sql"
)

func UserExists(db *sql.DB, username string) bool {
	var exists bool
	row := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM User_ WHERE username = $1)`, username)
	err := row.Scan(&exists)

	if err != nil {
		return false
	}
	return exists
}
