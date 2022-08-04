package app

import (
	"context"

	"github.com/mohammadmokh/writino/adaptor/email"
	"github.com/mohammadmokh/writino/adaptor/store/filesystem"
	mongodb "github.com/mohammadmokh/writino/adaptor/store/mongodb"
	"github.com/mohammadmokh/writino/adaptor/store/redis"
	"github.com/mohammadmokh/writino/config"
	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/interactor/auth"
	"github.com/mohammadmokh/writino/interactor/comment"
	"github.com/mohammadmokh/writino/interactor/post"
	"github.com/mohammadmokh/writino/interactor/user"
	"github.com/mohammadmokh/writino/interactor/verificationCode"
	"github.com/mohammadmokh/writino/jwt"
	"github.com/mohammadmokh/writino/validator"
)

type App struct {
	JwtSecret         string
	JwtParser         contract.ParseToken
	Auth              contract.AuthInteractor
	User              contract.UserInteractor
	Post              contract.PostInteractor
	Comment           contract.CommentInteractor
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
	image, err := filesystem.New(cfg.ImageFs)
	if err != nil {
		return App{}, err
	}

	mailService := email.New(cfg.Email)
	verficationCode := verificationCode.New(redisClient, verificationCode.Random, verificationCode.ParseVerificationTempl)
	auth := auth.New(MongoStore, []byte(cfg.JwtSecret), jwt.GenerateTokenPair, jwt.ParseRefToken)
	post := post.New(MongoStore)
	comment := comment.New(MongoStore)

	userBuilder := user.NewBuilder()
	userBuilder.
		SetCommentService(comment).
		SetMailService(mailService).
		SetPostSerivce(post).
		SetProfilePicStore(image).
		SetUserStore(MongoStore).
		SetVerCodeService(verficationCode)

	user := userBuilder.Build()

	return App{
		JwtSecret:         cfg.JwtSecret,
		JwtParser:         jwt.ParseToken,
		Auth:              auth,
		User:              user,
		Post:              post,
		Comment:           comment,
		CreatePostVal:     validator.ValidateCreatePost,
		UpdatePostVal:     validator.ValidateUpdatePost,
		RegisterVal:       validator.ValidateRegisterUser,
		UpdateUserVal:     validator.ValidateUpdateUser,
		UpdatePasswordVal: validator.ValidateUpdatePassword,
	}, nil
}
