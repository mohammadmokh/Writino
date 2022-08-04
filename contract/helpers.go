package contract

import "github.com/mohammadmokh/writino/entity"

type (
	Random                     func() (string, error)
	ParseVerificationEmailTmpl func(entity.VerificationCode) (string, error)
)
