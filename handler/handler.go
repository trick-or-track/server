package handler

import (
	"github.com/trick-or-track/server/data"
	"github.com/trick-or-track/server/user"
)

type Handler struct {
	dataStore data.Store
	userStore user.Store
}

func NewHandler(
	dataStore data.Store,
	userStore user.Store,
) *Handler {
	return &Handler{
		dataStore: dataStore,
		userStore: userStore,
	}
}
