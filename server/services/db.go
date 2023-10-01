package services

import (
  "database/sql"
  "log"
)

func ConnectDB(connStr string) *sql.DB {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }
  return db
}

