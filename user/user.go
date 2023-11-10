package user

import (
	"github.com/trick-or-track/server/model"
)

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}

type Store interface {
	GetByID(int) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
	Update(int, *CreateUserInput) error
	Delete(int) error
}
