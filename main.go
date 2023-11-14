package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/trick-or-track/server/db"
	"github.com/trick-or-track/server/handler"
	"github.com/trick-or-track/server/router"
	"github.com/trick-or-track/server/store"
)

func main() {
	r := router.New()
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "server running")
	})

	postgresDb, err := db.New()
	if err != nil {
		r.Logger.Fatal(err)
		return
	}

	v1 := r.Group("/api")
	dataStore := store.NewDataStore(postgresDb)
	userStore := store.NewUserStore(postgresDb)

	handler := handler.NewHandler(dataStore, userStore)
	handler.Register(v1)

	r.Logger.Fatal(r.Start(os.Getenv("PORT")))
}
