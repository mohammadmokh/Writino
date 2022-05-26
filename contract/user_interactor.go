package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/dto"
)

type UserInteractor interface {
	CheckUsername(context.Context, dto.CheckUsernameReq) (dto.CheckUsernameRes, error)
	CheckEmail(context.Context, dto.CheckEmailReq) (dto.CheckEmailRes, error)
	Register(context.Context, dto.RegisterReq) error
	Update(context.Context, dto.UpdateUserReq) error
	UpdatePassword(context.Context, dto.UpdatePasswordReq) error
	Find(context.Context, dto.FindUserReq) (dto.FindUserRes, error)
	DeleteAccount(context.Context, dto.DeleteUserReq) error
	VerifyUser(context.Context, dto.VerifyUserReq) error
}
