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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variable")
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL is not set in .env")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not set in .env, default fallback to 8080")
		port = "8080"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("✅ Connected to database successfully")

	userRepo := pg.NewUserRepo(db)
	_ = userRepo // "I know this is unused, shut up compiler"

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server starting on %s\n", addr)

	if err = http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
