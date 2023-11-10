package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	connStr := "user=postgres dbname=test_server password=postgrespw sslmode=disable port=55000"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
