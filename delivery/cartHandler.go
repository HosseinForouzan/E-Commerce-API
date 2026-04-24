package delivery

import (
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/labstack/echo/v4"
)

func (s Server) AddItem(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)
	var req param.AddItemRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	req.UserID = claims.UserID

	err := s.cartSvc.AddItem(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message" : "item created."})
}