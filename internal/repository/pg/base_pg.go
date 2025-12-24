package pg

import (
	"database/sql"
	"log"
)

func OpenDBEntity(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("[Database] Error opening database: ", err)
	}

	return db
}
