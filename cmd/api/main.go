package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	delivery "github.com/waldy-ctt/core-link-be/internal/delivery/http"
	"github.com/waldy-ctt/core-link-be/internal/repository/pg"
	authUC "github.com/waldy-ctt/core-link-be/internal/usecase/auth"
)

func main() {
	connStr, port := loadEnv()

	db := pg.OpenDBEntity(connStr)

	defer db.Close()

	checkDBHealth(db)

	// 1. Run DB Init / Migration
	if err := pg.RunMigrations(db); err != nil {
		log.Fatal("[Database] Failed to migrations: ", err)
	}

	// 2. Init repo
	userRepo := pg.NewUserRepo(db)
	_ = userRepo // "I know this is unused, shut up compiler"
	authRepo := pg.NewAuthRepo(db)
	_ = authRepo

	// 3. API Related Setting
	timeout := time.Duration(2) * time.Second // API Having 2 second timeout
	signupUC := authUC.NewSignupUseCase(userRepo, authRepo, timeout)

	authHandler := delivery.NewAuthHandler(signupUC)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", authHandler.Register)

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("[System] Server starting on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
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
		log.Fatal("[Database] Failed to connect to database: ", err)
	}
	fmt.Println("[Database] Connected to database successfully")
}
