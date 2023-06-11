package db

import (
	"database/sql"
	"log"
)

func ConnDb() *sql.DB {
	connStr := "postgres://<usuario>:<senha>@<host>/<banco>?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
