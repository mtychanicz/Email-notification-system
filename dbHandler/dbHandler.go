package dbHandler

import (
	"database/sql"
	"fmt"
	Models "golang-projects/dbModels"
	Sqlite "golang-projects/sqliteDB"
)

// DBHandler is the interface that defines the methods to interact with the database
type DBHandler interface {
	CreateIfDoesNotExist() error
	AddMessage(Content string, EventType string) (int, error)
	GetMessage(MessageId int) (*Models.Message, error)
}

type SQLiteHandler struct {
	DB *sql.DB
}

func (h *SQLiteHandler) CreateIfDoesNotExist() error {
	return Sqlite.CreateDB(h.DB)
}

func (h *SQLiteHandler) AddMessage(Content string, EventType string) (int, error) {
	fmt.Println("I'm running! @CreateUser @dbHandler")
	user := &Models.Message{Content: Content, EventType: EventType}
	return Sqlite.InsertMessage(h.DB, user)
}

func (h *SQLiteHandler) GetMessage(eventType string) (*Models.Message, error) {
	return Sqlite.GetMessage(h.DB, eventType)
}
