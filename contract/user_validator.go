package contract

import (
	"github.com/mohammadmokh/writino/dto"
)

type (
	ValidateRegisterUser   func(req dto.RegisterReq) error
	ValidateUpdateUser     func(req dto.UpdateUserReq) error
	ValidateUpdatePassword func(req dto.UpdatePasswordReq) error
)
