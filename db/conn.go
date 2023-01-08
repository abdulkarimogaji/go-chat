package db

import (
	"database/sql"
)

type Store struct {
	DB *sql.DB
}

// ConnectDatabase returns db.Store
func ConnectDatabase(uri string) (*Store, error) {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	// create tables
	// _, err = db.Exec(createTableUser)
	// if err != nil {
	// 	return nil, err
	// }

	return &Store{DB: db}, nil
}
