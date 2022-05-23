package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type VerficationCodeStore interface {
	FindVerCode(context.Context, string) (string, error)
	CreateVerCode(context.Context, entity.VerificationCode) error
}
