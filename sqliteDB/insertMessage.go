package sqliteDB

import (
	"database/sql"
	"fmt"
	Models "golang-projects/dbModels"

	_ "github.com/mattn/go-sqlite3"
)

func InsertMessage(db *sql.DB, message *Models.Message) (int, error) {
	// Open connection
	db, err := sql.Open("sqlite3", "messagesDB")
	if err != nil {
		return -1, err
	}
	defer db.Close()

	// Insert query
	sqlStatement := `
        INSERT INTO messages (content, eventType)
        VALUES (?, ?)`

	// Execute query
	result, err := db.Exec(sqlStatement, message.Content, message.EventType)
	if err != nil {
		return 0, err
	}

	// Get id of recently added Message
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Last added id here: %v", id)

	return int(id), nil
}
