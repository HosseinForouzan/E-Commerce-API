package delivery

import (
	"net/http"
	"strconv"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/labstack/echo/v4"
)

func (s Server) UserRegister(c echo.Context) error {
	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.UserSvc.Register(req)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) UserLogin(c echo.Context) error {
	var req param.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err:= s.UserSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) UserProfile(c echo.Context) error {
	var req param.ProfileRequest
	id := c.Param("id")
	the_id, _:= strconv.ParseUint(id, 10, 64)

	req.UserID =uint(the_id)
	resp, err := s.UserSvc.Profile(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)


}