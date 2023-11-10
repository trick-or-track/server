package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trick-or-track/server/router/middleware"
	"github.com/trick-or-track/server/utils"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleWare := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/user")
	guestUsers.GET("", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "hello")
	})
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleWare)
	user.GET("", h.CurrentUser)
	// user.PUT("", h.UpdateUser)

	// profiles := v1.Group("/profiles", jwtMiddleWare)
	// profiles.GET("/:username", h.GetProfile)

	data := v1.Group("/data", jwtMiddleWare)
	data.GET("", h.GetDataByUserID)
	data.POST("", h.AddData)
}
