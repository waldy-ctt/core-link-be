package pg

import "database/sql"

type PostgreRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *PostgreRepo {
	return &PostgreRepo {
		DB: db,
	}
}
