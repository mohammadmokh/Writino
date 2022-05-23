package contract

import "context"

type VerificationCodeInteractor interface {
	Find(context.Context, string) (string, error)
	Create(context.Context, string) (string, error)
}
