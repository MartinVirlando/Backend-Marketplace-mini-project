package middleware

import (
	"net/http"
	"strings"

	"backend/utils"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//Ambil heaeder authorization
			authHeader := c.Request().Header.Get("Authorization")

			//Cek Autheadernya ksong apa kaga
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
			}

			//Split "Bearer <token>" → ambil tokennya saja
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header format")
			}
			tokenString := parts[1]

			//Validasi tokennya
			claims, err := utils.ValidateToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			//Simpen Claims
			c.Set("user", claims)

			return next(c)

		}
	}

}
