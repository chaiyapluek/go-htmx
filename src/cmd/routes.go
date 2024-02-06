package main

import (
	"github.com/chaiyapluek/go-htmx/src/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Static("/static", "static")

	e.Use(middleware.Logger())

	e.GET("", handlers.HomeGetHandler)
	e.GET("/about", handlers.AboutGetHandler)
	e.GET("/contact", handlers.ContactGetHandler)
	e.GET("/reservation", handlers.ReservationGetHandler)

	return e
}
