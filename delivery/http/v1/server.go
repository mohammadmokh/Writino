package v1

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/gocastsian/writino/app"
	"gitlab.com/gocastsian/writino/config"
	"gitlab.com/gocastsian/writino/delivery/http/v1/auth"
	"gitlab.com/gocastsian/writino/delivery/http/v1/middleware"
	"gitlab.com/gocastsian/writino/delivery/http/v1/user"
)

type Server struct {
	server *echo.Echo
	cfg    config.ServerCfg
}

func New(app app.App, cfg config.ServerCfg) Server {

	e := echo.New()

	e.POST("/auth/login", auth.Login(app.Auth))
	e.POST("/auth/refresh", auth.Refresh(app.Auth))

	e.POST("/users", user.Register(app.User, app.RegisterVal))
	e.GET("/users/:username", user.Find(app.User))
	e.PATCH("/users", user.Update(app.User, app.UpdateUserVal),
		middleware.AuthMiddleware([]byte(app.JwtSecrete), app.JwtParser))
	e.DELETE("/users", user.Delete(app.User))

	e.POST("/check/username", user.CheckUsername(app.User))
	e.POST("/check/email", user.CheckEmail(app.User))

	e.PATCH("/update/password", user.UpdatePassword(app.User, app.UpdatePasswordVal))

	e.POST("/verify", user.Verify(app.User))

	return Server{
		server: e,
		cfg:    cfg,
	}
}

func (s Server) Run() error {

	return s.server.Start(s.cfg.Address)
}
