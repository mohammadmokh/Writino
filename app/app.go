package app

import (
	"context"

	"gitlab.com/gocastsian/writino/adaptor/email"
	mongodb "gitlab.com/gocastsian/writino/adaptor/store/mongodb"
	"gitlab.com/gocastsian/writino/adaptor/store/redis"
	"gitlab.com/gocastsian/writino/config"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/interactor/auth"
	"gitlab.com/gocastsian/writino/interactor/post"
	"gitlab.com/gocastsian/writino/interactor/user"
	"gitlab.com/gocastsian/writino/interactor/verificationCode"
	"gitlab.com/gocastsian/writino/jwt"
	"gitlab.com/gocastsian/writino/validator"
)

type App struct {
	JwtSecret         string
	JwtParser         contract.ParseToken
	Auth              contract.AuthInteractor
	User              contract.UserInteractor
	Post              contract.PostInteractor
	CreatePostVal     contract.ValidateCreatePost
	UpdatePostVal     contract.ValidateUpdatePost
	RegisterVal       contract.ValidateRegisterUser
	UpdateUserVal     contract.ValidateUpdateUser
	UpdatePasswordVal contract.ValidateUpdatePassword
}

func New(cfg config.Config) (App, error) {

	MongoStore, err := mongodb.New(context.TODO(), cfg.Mongo)
	if err != nil {
		return App{}, err
	}
	redisClient, err := redis.New(context.TODO(), cfg.Redis)
	if err != nil {
		return App{}, err
	}

	mailService := email.New(cfg.Email)
	verficationCode := verificationCode.New(redisClient, verificationCode.Random, verificationCode.ParseVerificationTempl)
	auth := auth.New(MongoStore, []byte(cfg.JwtSecret), jwt.GenerateTokenPair, jwt.ParseRefToken)
	user := user.New(MongoStore, mailService, verficationCode)
	post := post.New(MongoStore)

	return App{
		JwtSecret:         cfg.JwtSecret,
		JwtParser:         jwt.ParseToken,
		Auth:              auth,
		User:              user,
		Post:              post,
		CreatePostVal:     validator.ValidateCreatePost,
		UpdatePostVal:     validator.ValidateUpdatePost,
		RegisterVal:       validator.ValidateRegisterUser,
		UpdateUserVal:     validator.ValidateUpdateUser,
		UpdatePasswordVal: validator.ValidateUpdatePassword,
	}, nil
}
