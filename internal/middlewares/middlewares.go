package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const desiredRole = "admin"

// FindAdmin проверяет header запроса на наличии desiredRole
func FindAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get("User-Role")
		if strings.EqualFold(desiredRole, userRole) {
			log.Println("red button user detected")
		}
		return next(c)
	}
}
