package contract

import "gitlab.com/gocastsian/writino/entity"

type (
	VerificationCodeGen    func() (string, error)
	ParseVerificationTempl func(user entity.User) (string, error)
)
