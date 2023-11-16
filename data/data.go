package data

import "github.com/trick-or-track/server/model"

type Store interface {
	Create(*model.Data) error
	GetByUserID(int, int, int) ([]*model.Data, error)
	GetYearly(int, int) ([]*model.Data, error)
}
