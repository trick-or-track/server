package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/trick-or-track/server/model"
	"github.com/trick-or-track/server/utils"
)

// type userUpdateRequest struct {
// 	User struct {
// 		UserName string `json:"username"`
// 		Emal string `json:"username"`
// 		UserName string `json:"username"`
// 	}
// }

type UserRegistrationRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *UserRegistrationRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = utils.Format(r.User.Username)
	u.Email = utils.Format(r.User.Email)
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = utils.Format(r.User.Email)
	u.Password = r.User.Password
	return nil
}

type DataAddRequest struct {
	Data struct {
		Year  int `json:"year"`
		One   int `json:"one"`
		Two   int `json:"two"`
		Three int `json:"three"`
		Four  int `json:"four"`
		Five  int `json:"five"`
		Six   int `json:"six"`
		Seven int `json:"seven"`
		Eight int `json:"eight"`
		Nine  int `json:"nine"`
		Ten   int `json:"ten"`
	} `json:"data"`
}

func (r *DataAddRequest) bind(c echo.Context, d *model.Data) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	d.UserID = userIDFromToken(c)
	d.Year = r.Data.Year
	d.One = r.Data.One
	d.Two = r.Data.Two
	d.Three = r.Data.Three
	d.Four = r.Data.Four
	d.Five = r.Data.Five
	d.Six = r.Data.Six
	d.Seven = r.Data.Seven
	d.Eight = r.Data.Eight
	d.Nine = r.Data.Nine
	d.Ten = r.Data.Ten
	return nil
}

// #region helper methods

func userIDFromToken(c echo.Context) int {
	id, ok := c.Get("user").(int)
	if !ok {
		return 0
	}
	return id
}

// #endregion helper methods
