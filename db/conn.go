package db

import (
	"database/sql"
	"fmt"
)

type Store struct {
	DB *sql.DB
}

// ConnectDatabase returns db.Store
func ConnectDatabase(uri string) (*Store, error) {
	conn, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, fmt.Errorf("Error Connecting %v", err)
	}

	// create tables
	_, err = conn.Exec(createTableUser)
	if err != nil {
		return nil, fmt.Errorf("Error Creating table user %s", err)
	}
	_, err = conn.Exec(createTableRoom)
	if err != nil {
		return nil, fmt.Errorf("Error Creating table room %s", err)
	}
	_, err = conn.Exec(createTableChat)
	if err != nil {
		return nil, fmt.Errorf("Error Creating table chat %s", err)
	}

	return &Store{DB: conn}, nil
}
