package contract

import (
	"gitlab.com/gocastsian/writino/dto"
)

type (
	ValidateRegisterUser   func(req dto.RegisterReq) error
	ValidateUpdateUser     func(req dto.UpdateUserReq) error
	ValidateUpdatePassword func(req dto.UpdatePasswordReq) error
)
