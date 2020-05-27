package dbhelper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() *sql.DB {
	database, err := sql.Open("sqlite3", "./crebit.db")
	if err != nil {
		log.Fatal(err)
	}
	return database
}
