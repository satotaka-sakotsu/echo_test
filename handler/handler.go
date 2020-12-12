package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
