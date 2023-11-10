package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trick-or-track/server/model"
	"github.com/trick-or-track/server/utils"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &UserRegistrationRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	err := h.userStore.Create(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (h *Handler) Login(c echo.Context) error {
	var u model.User
	req := &userLoginRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	foundUser, err := h.userStore.GetByEmail(u.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewError(err))
	}

	if ok := foundUser.CheckPassword(u.Password); !ok {
		return c.JSON(http.StatusUnauthorized, utils.NewError(fmt.Errorf("invalid credentials")))
	}

	return c.JSON(http.StatusCreated, newUserResponse(foundUser))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newUserResponse(u))
}
