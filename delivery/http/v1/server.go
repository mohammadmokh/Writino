package v1

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/gocastsian/writino/app"
	"gitlab.com/gocastsian/writino/config"
	"gitlab.com/gocastsian/writino/delivery/http/v1/auth"
	"gitlab.com/gocastsian/writino/delivery/http/v1/comment"
	"gitlab.com/gocastsian/writino/delivery/http/v1/middleware"
	"gitlab.com/gocastsian/writino/delivery/http/v1/post"
	"gitlab.com/gocastsian/writino/delivery/http/v1/user"
)

type Server struct {
	server *echo.Echo
	cfg    config.ServerCfg
}

func New(app app.App, cfg config.Config) Server {

	e := echo.New()

	e.Use(middleware.AuthMiddleware([]byte(app.JwtSecret), app.JwtParser))

	e.POST("/auth/login", auth.Login(app.Auth))
	e.POST("/auth/refresh", auth.Refresh(app.Auth))

	e.POST("/users", user.Register(app.User, app.RegisterVal))
	e.GET("/users/:id", user.Find(app.User, cfg.Server))
	e.PATCH("/users", user.Update(app.User, app.UpdateUserVal))
	e.DELETE("/users", user.Delete(app.User))
	e.POST("/check/username", user.CheckUsername(app.User))
	e.POST("/check/email", user.CheckEmail(app.User))
	e.PATCH("/update/password", user.UpdatePassword(app.User, app.UpdatePasswordVal))
	e.POST("/verify", user.Verify(app.User))
	e.PATCH("/update/avatar", user.UpdateAvatar(app.User, cfg.Server))

	e.POST("/posts", post.CreatePost(app.Post, app.CreatePostVal, cfg.Server))
	e.GET("/posts/:id", post.FindPostByID(app.Post, cfg.Server))
	e.PATCH("/posts/:id", post.UpdatePost(app.Post, app.UpdatePostVal))
	e.DELETE("/posts/:id", post.Delete(app.Post))
	e.GET("/posts/search", post.SearchPost(app.Post, cfg.Server))
	e.GET("/users/:id/posts", post.FindUsersPosts(app.Post, cfg.Server))
	e.GET("/posts", post.FindAll(app.Post, cfg.Server))
	e.GET("/posts", post.FindAll(app.Post, cfg.Server))
	e.PUT("/posts/:id/like", post.LikePost(app.Post))

	e.POST("/posts/:id/comments", comment.CreateComment(app.Comment))
	e.GET("/posts/:id/comments", comment.FindCommentsByPostID(app.Comment, cfg.Server))

	e.Static("/images", cfg.ImageFs.BasePath)

	return Server{
		server: e,
		cfg:    cfg.Server,
	}
}

func (s Server) Run() error {

	return s.server.Start(s.cfg.Address)
}
