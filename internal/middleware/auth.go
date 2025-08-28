package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func TokenAuthMiddleware(expectedToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization header"})
			}

			// Expect "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid auth header format"})
			}

			token := parts[1]
			if token != expectedToken {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Token is valid, continue
			return next(c)
		}
	}
}
