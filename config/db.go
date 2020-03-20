package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Can't open the db. err=%v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error while pinging the DB. err=%v", err)
	}

	log.Println("Database connected.")
}
