package sqliteDB

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//TODO: Look into orm module for Go

// This won't remove existing db if run again
func CreateDB(db *sql.DB) error {
	// Open connection
	db, err := sql.Open("sqlite3", "usersDB")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create Users table
	messageTable := `
	CREATE TABLE IF NOT EXISTS users (
		MessageId INTEGER PRIMARY KEY AUTOINCREMENT,
		Content TEXT NOT NULL,
		EventType TEXT NOT NULL
	);
	`
	_, err = db.Exec(messageTable)
	if err != nil {
		return err
	}
	fmt.Println("MessageTable created successfully")
	return nil
}
