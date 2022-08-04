package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadmokh/writino/contract"
)

const CtxUserKey = "user"

func AuthMiddleware(secret []byte, parser contract.ParseToken) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				c.Set(CtxUserKey, nil)
				return next(c)
			}
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 {
				c.Set(CtxUserKey, nil)
				return next(c)
			}
			if headerParts[0] != "Bearer" {
				c.Set(CtxUserKey, nil)
				return next(c)
			}
			user, err := parser(secret, headerParts[1])
			if err != nil {
				c.Set(CtxUserKey, nil)
				return next(c)
			}
			c.Set(CtxUserKey, user)
			return next(c)
		}
	}
}
