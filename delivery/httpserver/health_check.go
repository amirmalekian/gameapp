package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) healthCheck(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.JSON(http.StatusOK, echo.Map{
		"message": "ok",
	})
}
