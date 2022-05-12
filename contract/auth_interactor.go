package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/dto"
)

type AuthInteractor interface {
	Login(context.Context, dto.LoginReq) (dto.LoginResponse, error)
	RefreshToken(context.Context, dto.RefreshReq) (dto.RefreshResponse, error)
}
