package server

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

//opening database
func OpenDB(configStr string) error {
	var err error
	db, err = sqlx.Open("postgres", configStr)
	return err
}

//closing database
func CloseDB() error {
	return db.Close()
}
