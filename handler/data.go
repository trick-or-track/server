package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trick-or-track/server/model"
	"github.com/trick-or-track/server/utils"
)

func (h *Handler) AddData(c echo.Context) error {
	var d model.Data
	req := &DataAddRequest{}
	if err := req.bind(c, &d); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if err := h.dataStore.Create(&d); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newDataResponse(append([]*model.Data{}, &d)))
}

func (h *Handler) GetDataByUserID(c echo.Context) error {
	userID := userIDFromToken(c)
	_from := c.QueryParam("from")
	_to := c.QueryParam("to")

	from, err := strconv.Atoi(_from)
	if err != nil {
		from = utils.YearsAgo(5)
	}
	to, err := strconv.Atoi(_to)
	if err != nil {
		to = utils.CurrentYear()
	}
	d, err := h.dataStore.GetByUserID(userID, from, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newDataResponse(d))
}

func (h *Handler) GetDataYearly(c echo.Context) error {
	_from := c.QueryParam("from")
	_to := c.QueryParam("to")

	from, err := strconv.Atoi(_from)
	if err != nil {
		from = utils.YearsAgo(5)
	}
	to, err := strconv.Atoi(_to)
	if err != nil {
		to = utils.CurrentYear()
	}

	d, err := h.dataStore.GetYearly(from, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newDataResponse(d))
}
