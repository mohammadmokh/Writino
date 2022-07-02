package post

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

func CreatePost(i contract.PostInteractor, validator contract.ValidateCreatePost, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.CreatePostReq{}
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}
		user := userCtx.(entity.User)
		req.AuthorID = user.Id

		validationErrs := validator(req)
		if validationErrs != nil {
			return c.JSON(http.StatusBadRequest, validationErrs)
		}

		res, err := i.CreatePost(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		//return link of resource
		res.Post = cfg.Address + "/posts/" + res.Post
		return c.JSON(http.StatusCreated, res)
	}
}

func FindPostByID(i contract.PostInteractor, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.FindPostByIDReq{}
		id := c.Param("id")
		userCtx := c.Get(middleware.CtxUserKey)
		user := userCtx.(entity.User)

		req.ID = id
		req.UserID = user.Id

		post, err := i.FindPostByID(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		// send link of user
		post.Author = cfg.Address + "/users/" + post.Author
		return c.JSON(http.StatusOK, post)
	}
}

func UpdatePost(i contract.PostInteractor, validator contract.ValidateUpdatePost) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.UpdatePostReq{}
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}
		user := userCtx.(entity.User)
		req.AuthorID = user.Id
		req.Id = c.Param("id")

		validationErrs := validator(req)
		if validationErrs != nil {
			return c.JSON(http.StatusBadRequest, validationErrs)
		}

		err = i.UpdatePost(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"msg": "post updated"})
	}
}

func SearchPost(i contract.PostInteractor, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		var err error
		req := dto.SearchPostReq{}

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

		req.Query = c.QueryParam("query")

		posts, err := i.SearchPost(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusOK, posts)
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		for i := 0; i < len(posts.Posts); i++ {
			// send link of resources
			posts.Posts[i].Author = cfg.Address + "/users/" + posts.Posts[i].Author
			posts.Posts[i].Link = cfg.Address + "/posts/" + posts.Posts[i].ID
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func FindUsersPosts(i contract.PostInteractor, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		var err error
		req := dto.FindUsersPostsReq{}

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

		req.UserID = c.Param("id")

		posts, err := i.FindUsersPosts(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusOK, posts)
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		for i := 0; i < len(posts.Posts); i++ {
			// send link of resources
			posts.Posts[i].Link = cfg.Address + "/posts/" + posts.Posts[i].ID
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func Delete(i contract.PostInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.DeletePostReq{}
		req.ID = c.Param("id")
		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}
		user := userCtx.(entity.User)
		req.UserID = user.Id

		err := i.DeletePost(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusBadRequest, echo.Map{"errors": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"errors": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "post deleted"})
	}
}

func FindAll(i contract.PostInteractor, cfg config.ServerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		var err error
		req := dto.SearchPostReq{}

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

		if len(c.QueryParam("query")) == 0 {
			req.Query = "newest"
		} else {
			req.Query = c.QueryParam("query")
		}

		posts, err := i.FindAll(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusOK, posts)
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		for i := 0; i < len(posts.Posts); i++ {
			// send link of resources
			posts.Posts[i].Author = cfg.Address + "/users/" + posts.Posts[i].Author
			posts.Posts[i].Link = cfg.Address + "/posts/" + posts.Posts[i].ID
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func LikePost(i contract.PostInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.LikePostReq{}
		req.PostID = c.Param("id")
		userCtx := c.Get(middleware.CtxUserKey)
		if userCtx == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": errors.ErrInvalidToken.Error()})
		}
		user := userCtx.(entity.User)
		req.UserID = user.Id

		err := i.LikePost(c.Request().Context(), req)
		if err != nil {
			if err == errors.ErrNotFound {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"msg": "post liked"})
	}
}
