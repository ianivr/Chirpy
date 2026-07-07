package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ianivr/chirpy/internal/database"
	"github.com/ianivr/chirpy/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("JWT_SECRET")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	log.Fatal(server.Start(dbQueries, platform, secret))
}
