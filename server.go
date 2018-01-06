package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "query, 200 OK")
	})
	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, getUser(id))
	})
	e.Logger.Fatal(e.Start(":1234"))
}

func getUser(id int) *User {
	name := "myname"
	u := &User{
		Id:   id,
		Name: name,
	}
	return u
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
