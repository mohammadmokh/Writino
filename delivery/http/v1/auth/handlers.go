package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/dto"
	"github.com/mohammadmokh/writino/errors"
)

func Login(i contract.AuthInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.LoginReq{}
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		tokenPair, err := i.Login(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrInvalidCredentials {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, tokenPair)
	}
}

func Refresh(i contract.AuthInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {
		req := dto.RefreshReq{}
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		tokenPair, err := i.RefreshToken(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrInvalidToken {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, tokenPair)
	}
}
