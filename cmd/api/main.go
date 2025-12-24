package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/waldy-ctt/core-link-be/internal/repository/pg"
)

func main() {
	connStr, port := loadEnv()

	db := pg.OpenDBEntity(connStr)

	defer db.Close()

	checkDBHealth(db)

	if err := pg.RunMigrations(db); err != nil {
		log.Fatal("[Database] Failed to migrations: ", err)
	}

	userRepo := pg.NewUserRepo(db)
	_ = userRepo // "I know this is unused, shut up compiler"

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("[System] Server starting on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() (string, string) {
	if err := godotenv.Load(); err != nil {
		log.Println("[System] No .env file found, relying on system environment variable")
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("[System] DB_URL is not set in .env")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("[System] SERVER_PORT is not set in .env, default fallback to 8080")
		port = "8080"
	}

	return connStr, port
}

func checkDBHealth(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("[Database] Connected to database successfully")
}
