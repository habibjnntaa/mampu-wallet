package database

import (
	"database/sql"
	"log"
)

func SeedData(db *sql.DB) {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM users)").Scan(&exists)
	if exists {
		log.Println("Data sudah tersedia, skip seeding.")
		return
	}

	// Menggunakan QueryRow untuk mendapatkan ID yang baru dibuat (RETURNING id)
	var lastID int64
	err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", "example", "example@example.com").Scan(&lastID)
	if err == nil {
		db.Exec("INSERT INTO wallets (user_id, balance) VALUES ($1, $2)", lastID, 1000000)
	}

	log.Println("Seeding Data berhasil.")
}
