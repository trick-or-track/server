package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/trick-or-track/server/model"
	"github.com/trick-or-track/server/user"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id int) (*model.User, error) {
	var m model.User
	if err := us.db.QueryRow(`
	SELECT id, username, email FROM users 
	WHERE id = $1
	AND is_deleted = false`, id).Scan(
		&m.ID,
		&m.Username,
		&m.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with id %d not found", id)
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByEmail(email string) (*model.User, error) {
	var m model.User
	if err := us.db.QueryRow(`
	SELECT id, username, email, password FROM users 
	WHERE email = $1
	AND is_deleted = false`, email).Scan(
		&m.ID,
		&m.Username,
		&m.Email,
		&m.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with email %v not found", email)
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(user *model.User) error {
	query, err := us.db.Prepare(`
	INSERT INTO users (
		email, username, password, created_at
	) VALUES ($1, $2, $3, $4) 
	RETURNING id;`)
	if err != nil {
		return err
	}
	var insertedId int
	if err := query.QueryRow(
		user.Email,
		user.Username,
		user.Password,
		time.Now().UTC()).Scan(&insertedId); err != nil {

		if err.Error() == "pq: duplicate key value violates unique constraint \"user_unique_username\"" {
			return fmt.Errorf("Username taken")
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"user_unique_email\"" {
			return fmt.Errorf("Email taken")
		}
		return err
	}
	user.ID = insertedId
	return nil
}

func (us *UserStore) Update(id int, userInput *user.CreateUserInput) error {
	query, err := us.db.Prepare(`
	UPDATE users SET 
		email=$1, 
		username=$2;`)
	if err != nil {
		return err
	}
	if _, err := query.Exec(
		userInput.Email,
		userInput.Username); err != nil {
		return err
	}
	return nil
}

func (us *UserStore) Delete(id int) error {
	query, err := us.db.Prepare(`
	UPDATE users SET is_deleted=true
	WHERE id = $1;`)
	if err != nil {
		return err
	}
	if _, err := query.Exec(id); err != nil {
		return err
	}
	return nil
}
