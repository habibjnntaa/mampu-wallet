package database

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id BIGSERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS wallets (
            id BIGSERIAL PRIMARY KEY,
            user_id BIGINT NOT NULL UNIQUE,
            balance NUMERIC(15, 2) DEFAULT 0.00,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
        );`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatalf("Gagal migrasi Postgres: %v", err)
		}
	}

	log.Println("Migrasi Table berhasil.")
}
