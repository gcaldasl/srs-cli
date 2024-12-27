package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) (*SQLiteRepository, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(homeDir, ".srs.db")
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := initDB(db); err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

func initDB(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS cards (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        front_side TEXT NOT NULL,
        back_side TEXT NOT NULL,
        last_reviewed DATETIME,
        next_review DATETIME,
        interval INTEGER,
        ease_factor REAL
    );`

	_, err := db.Exec(query)
	return err
}
