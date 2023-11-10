package data

import "github.com/trick-or-track/server/model"

type Store interface {
	Create(*model.Data) error
	GetByUserID(int) ([]*model.Data, error)
}
