package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/dto"
)

type UserInteractor interface {
	CheckUsername(context.Context, dto.CheckUsernameReq) (dto.CheckUsernameRes, error)
	CheckEmail(context.Context, dto.CheckEmailReq) (dto.CheckEmailRes, error)
	Register(context.Context, dto.RegisterReq, ValidateRegisterUser) error
	Update(context.Context, dto.UpdateUserReq, ValidateUpdateUser) error
	UpdatePassword(context.Context, dto.UpdatePasswordReq, ValidateUpdatePassword) error
	Find(context.Context, dto.FindUserReq) (dto.FindUserRes, error)
	DeleteAccount(context.Context, dto.DeleteUserReq) error
	VerifyUser(context.Context, dto.VerifyUserReq) error
}
