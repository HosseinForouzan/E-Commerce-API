package delivery

import (
	"net/http"
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice/authservice"
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
	authSvc := authservice.New("secret", "at", "rt", time.Hour * 24, time.Hour * 24 * 7)
	authToken := c.Request().Header.Get("Authorization")
	claims, err := authSvc.ParseToken(authToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "token is not valid.")
	}
	
	resp, err := s.UserSvc.Profile(param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)


}