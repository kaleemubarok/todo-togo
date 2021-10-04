package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewSqliteConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./db/todotogo.db")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
