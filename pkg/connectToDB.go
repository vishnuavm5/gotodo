package pkg

import (
	"database/sql"
	"hellocheck/internal/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDataBase() *database.Queries {
	godotenv.Load()

	conString := os.Getenv("DB_URL")
	if conString == "" {
		log.Fatal("No connection string")
	}

	conn, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(conn)
	return queries

}
