package pg

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	queries := []string{
		// 1. Users Table
		`CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			display_name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		);`,
		
		// 2. Auth Table (Links to Users)
		`CREATE TABLE IF NOT EXISTS auth (
			user_id TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			refresh_token TEXT,
			last_login TIMESTAMP
		);`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return fmt.Errorf("[Database] migration failed: %w", err)
		}
	}
	
	fmt.Println("[Database] Database tables initialized (or already existed)")
	return nil
}
