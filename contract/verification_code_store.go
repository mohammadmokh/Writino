package contract

import (
	"context"

	"github.com/mohammadmokh/writino/entity"
)

type VerficationCodeStore interface {
	FindVerCode(context.Context, string) (string, error)
	CreateVerCode(context.Context, entity.VerificationCode) error
}
