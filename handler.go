package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	Handler struct {
		Db map[string]string
	}

	migrationResultJSON struct {
		Id string `json:"Id"`
	}
)

func (h Handler) GetUser(c echo.Context) error {
	user := h.Db["1"]
	if user == "" {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, migrationResultJSON{Id: user})
}
