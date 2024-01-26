package pkg

import (
	"database/sql"
	"hellocheck/internal/database"
	"log"
	"os"
)

var conn *sql.DB

func InitDB() {
	// godotenv.Load()

	conString := os.Getenv("DB_URL")
	if conString == "" {
		log.Fatal("No connection string")
	}
	var err error

	conn, err = sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectToDataBase() *database.Queries {
	return database.New(conn)
}
