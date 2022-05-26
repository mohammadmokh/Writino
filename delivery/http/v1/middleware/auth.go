package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/errors"
)

const CtxUserKey = "user"

func AuthMiddleware(secret []byte, parser contract.ParseToken) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.NoContent(http.StatusUnauthorized)
			}
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 {
				return c.NoContent(http.StatusUnauthorized)
			}
			if headerParts[0] != "Bearer" {
				return c.NoContent(http.StatusUnauthorized)
			}
			user, err := parser(secret, headerParts[1])
			if err != nil {
				status := http.StatusInternalServerError
				if err == errors.ErrInvalidToken {
					status = http.StatusUnauthorized
				}

				return c.NoContent(status)
			}

			c.Set(CtxUserKey, user)
			return next(c)
		}
	}
}
