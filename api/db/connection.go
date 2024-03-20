package db

import (
	"database/sql"
)

type Connection struct {
	dns string
	DB  *sql.DB
}

func MakeDBConnectionFrom(dsn string) (Connection, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return Connection{}, err
	}

	return Connection{
		dns: dsn,
		DB:  db,
	}, nil
}

func (receiver Connection) Close() error {
	err := receiver.DB.Close()

	if err != nil {
		return err
	}

	return nil
}
