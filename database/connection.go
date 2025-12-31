package database

import (
	"database/sql"
	"fmt"
	"log"
	"mampu-wallet/internal/tools"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	tools.LoadEnv()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
