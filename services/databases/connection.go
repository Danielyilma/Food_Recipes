package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbConnect() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	HOST := os.Getenv("HOST")

	url := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		POSTGRES_USER, POSTGRES_DB, POSTGRES_PASSWORD, HOST)
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return nil
	} else {
		log.Print("Successfully Connected")
	}

	return db
}

var DBConnection *sqlx.DB = DbConnect()
