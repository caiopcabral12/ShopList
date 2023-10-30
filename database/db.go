package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection := "user=postgres dbname=shop password=0102 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
