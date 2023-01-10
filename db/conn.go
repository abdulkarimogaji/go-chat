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
	prep, err := conn.Prepare(createTableUser)
	if err != nil {
		return nil, fmt.Errorf("Error Preparing %v", err)
	}
	_, err = prep.Exec()
	if err != nil {
		return nil, fmt.Errorf("Error Executing %v", err)
	}

	return &Store{DB: conn}, nil
}
