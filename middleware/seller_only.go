package middleware

import (
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SellerOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*utils.JwtCustomClaims)
			if !ok || user.Role != "seller" {
				return utils.ErrorResponse(c, http.StatusForbidden, "Access Denied")
			}
			return next(c)
		}
	}

}
