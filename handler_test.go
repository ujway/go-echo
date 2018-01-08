package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

var (
	testMockDB = map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")
	h := Handler{testMockDB}

	h.GetUser(c)
	if rec.Body.String() != "{\"id\":\"one\"}" {
		t.Errorf("expected response Id: one, got %s", rec.Body.String())
	}
}
