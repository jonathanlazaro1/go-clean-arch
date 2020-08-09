package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	//Postgres Go driver
	_ "github.com/lib/pq"
)

// NewPGSQLDB creates a new instance of PGSQL DB Connection
func NewPGSQLDB() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading environment config")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PGSQL!")
	return db
}
