package main

import (
	"myproject/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

var (
	mockDB = map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}
)

func main() {
	e := echo.New()
	h := handler.Handler{mockDB}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", h.GetUser)

	e.Run(standard.New(":1323"))
}
