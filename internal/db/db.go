package db

import "database/sql"

func Connect(DatabaseUrl string) (*sql.DB, error) {
	connect, err := sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return connect, nil

}
