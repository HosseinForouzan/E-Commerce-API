package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			role, ok := c.Get("claims").(string)
			fmt.Println(role)
			fmt.Println(c.Request().Header)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}

			for _, allowed := range roles {
				if role == allowed {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "forbidden",
			})
		}
	}
}