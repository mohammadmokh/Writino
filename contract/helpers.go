package contract

import "gitlab.com/gocastsian/writino/entity"

type (
	Random                     func() (string, error)
	ParseVerificationEmailTmpl func(entity.VerificationCode) (string, error)
)
