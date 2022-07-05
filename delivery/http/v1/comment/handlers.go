package comment

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.com/gocastsian/writino/config"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/delivery/http/v1/middleware"
	"gitlab.com/gocastsian/writino/dto"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
)

func CreateComment(i contract.CommentInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.CreateCommentReq{}
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}
		user := userCtx.(entity.User)
		req.UserID = user.Id
		req.PostID = c.Param("id")

		err = i.CreateComment(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}

func FindCommentsByPostID(i contract.CommentInteractor, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		var err error
		req := dto.FindCommentReq{}

		// set default values for pagination
		if len(c.QueryParam("limit")) == 0 {
			req.Limit = 5
		} else {
			req.Limit, err = strconv.Atoi(c.QueryParam("limit"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
		}

		if len(c.QueryParam("page")) == 0 {
			req.Page = 1
		} else {
			req.Page, err = strconv.Atoi(c.QueryParam("page"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
		}

		req.PostID = c.Param("id")

		comments, err := i.FindCommentsByPostID(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusOK, comments)
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		for i := 0; i < len(comments.Comments); i++ {
			// send link of resources
			comments.Comments[i].User = cfg.Address + "/users/" + comments.Comments[i].User
		}

		return c.JSON(http.StatusOK, comments)
	}
}
