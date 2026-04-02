package middleware

import (
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//Ambil customclaim dari context
			user, ok := c.Get("user").(*utils.JwtCustomClaims)

			if !ok || user.Role != "admin" {
				return utils.ErrorResponse(c, http.StatusForbidden, "Access Denied")
			}
			return next(c)
		}

	}
}
