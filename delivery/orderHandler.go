package delivery

import (
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/labstack/echo/v4"
)

func (s Server) Checkout(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)

	resp, err := s.orderSvc.Checkout(c.Request().Context(), claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)
}