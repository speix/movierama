package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sqlx.DB
}

func NewDB(path string) *Database {

	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{DB: db}
}
