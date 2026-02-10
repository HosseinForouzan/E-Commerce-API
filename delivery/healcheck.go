package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Service is OK!")
}