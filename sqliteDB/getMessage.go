package sqliteDB

import (
	"database/sql"
	Models "golang-projects/dbModels"

	_ "github.com/mattn/go-sqlite3"
)

func GetMessage(db *sql.DB, eventType string) (*Models.Message, error) {
	// Open connection to db
	db, err := sql.Open("sqlite3", "usersDB")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Get user
	var message Models.Message
	row := db.QueryRow("SELECT * FROM users WHERE EventType = ?", eventType)
	err = row.Scan(&message.MessageId, &message.Content, &message.EventType)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
