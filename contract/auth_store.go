package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type AuthStore interface {
	FindUser(context.Context, string) (entity.User, error)
	FindUserByEmail(context.Context, string) (entity.User, error)
}
