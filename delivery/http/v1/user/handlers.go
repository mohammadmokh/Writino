package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/delivery/http/v1/middleware"
	"gitlab.com/gocastsian/writino/dto"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
)

func Register(i contract.UserInteractor, validator contract.ValidateRegisterUser) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.RegisterReq{}
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		validationErrs := validator(req)
		if validationErrs != nil {
			return c.JSON(http.StatusBadRequest, validationErrs)
		}

		err = i.Register(c.Request().Context(), req)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, echo.Map{"msg": "user created"})
	}
}

func CheckUsername(i contract.UserInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.CheckUsernameReq{}
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		res, err := i.CheckUsername(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func CheckEmail(i contract.UserInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.CheckEmailReq{}
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		res, err := i.CheckEmail(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func Update(i contract.UserInteractor, validator contract.ValidateUpdateUser) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.UpdateUserReq{}

		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}

		user := userCtx.(entity.User)
		req.ID = user.Id

		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		validationErrs := validator(req)
		if validationErrs != nil {
			return c.JSON(http.StatusBadRequest, validationErrs)
		}

		err = i.Update(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "user updated"})
	}
}

func Delete(i contract.UserInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.DeleteUserReq{}

		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}

		user := userCtx.(entity.User)
		req.Id = user.Id

		err := i.DeleteAccount(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "user deleted"})
	}
}

func Find(i contract.UserInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.FindUserReq{}

		username := c.Param("username")
		req.Username = username

		res, err := i.Find(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func UpdatePassword(i contract.UserInteractor, validator contract.ValidateUpdatePassword) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.UpdatePasswordReq{}

		user := c.Get(middleware.CtxUserKey).(entity.User)
		req.ID = user.Id

		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		validationErrs := validator(req)
		if validationErrs != nil {
			return c.JSON(http.StatusBadRequest, validationErrs)
		}

		err = i.UpdatePassword(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrInvalidCredentials {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "password updated"})
	}
}

func Verify(i contract.UserInteractor) echo.HandlerFunc {

	return func(c echo.Context) error {

		req := dto.VerifyUserReq{}
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		err = i.VerifyUser(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrInvalidCredentials {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "user verified"})
	}
}
