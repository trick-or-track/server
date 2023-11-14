package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Toronto",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	seed := `
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	is_deleted BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS user_unique_email ON users(email);
CREATE UNIQUE INDEX IF NOT EXISTS user_unique_username ON users(username);

CREATE TABLE IF NOT EXISTS data(
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	year INT NOT NULL,
	one INT DEFAULT 0,
	two INT DEFAULT 0,
	three INT DEFAULT 0,
	four INT DEFAULT 0,
	five INT DEFAULT 0,
	six INT DEFAULT 0,
	seven INT DEFAULT 0,
	eight INT DEFAULT 0,
	nine INT DEFAULT 0,
	ten INT DEFAULT 0,
	created_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS data_unique_year_user_id ON data(year,user_id);
`

	_, err := db.Exec(seed)
	if err != nil {
		panic(err)
	}

	return nil
}
