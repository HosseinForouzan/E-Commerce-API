package middleware

import (
	"fmt"
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/service/authorizationservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/labstack/echo/v4"
)

func AccessCheck(service authorizationservice.Service) echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("claims").(*authservice.Claims)
			fmt.Println("salam", claims)
			isAdmin, err := service.CheckAccess(uint8(claims.UserID))
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "error in getting user status access.",
				})
			}

			if !isAdmin {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": "you don't have access.",
				})
			}

			return next(c)
		}

	}
}