package middleware

import (
	"eventhub/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Отсутствует токен"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		username, err := utils.ValidateJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Невалидный токен"})
		}

		c.Set("username", username)
		return next(c)
	}
}
