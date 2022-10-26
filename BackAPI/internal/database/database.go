package database

import (
	"database/sql"
)

var db *sql.DB

func OpenBD(configStr string) error {
	var err error
	db, err = sql.Open("postgres", configStr)
	return err
}

func CloseDB() error {
	return db.Close()
}
