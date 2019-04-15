package contracts

import (
	"database/sql"
	"log"

	
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3306)/kumparannews")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
