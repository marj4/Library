package db

import (
	"database/sql"
	"fmt"
)

const (
	ConectError = "Cant connect to database,error: %v"
	PingError   = "Cant ping database,%v"
)

func Connect(DatabaseUrl string) (*sql.DB, error) {
	connect, err := sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf(ConectError, err)
	}
	return connect, nil
}

func Ping(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf(PingError, err)
	}
	return nil
}
