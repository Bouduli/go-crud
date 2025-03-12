package db

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = nil

//TODO: figure out some way to handle this "better"

// Starts the SQLite database and returns it
func Start() (*sql.DB, error) {
	appRoot, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path.Join(appRoot, "db", "app.db"))
	if err != nil {
		return nil, err
	}

	DB = db

	return DB, nil
}

// Close
// closes the database
func Close() error {
	return DB.Close()
}
