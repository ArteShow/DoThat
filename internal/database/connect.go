package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB
}

func Connect(path string) (*Database, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
