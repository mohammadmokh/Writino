package contract

import (
	"context"

	"github.com/mohammadmokh/writino/entity"
)

type AuthStore interface {
	FindUser(context.Context, string) (entity.User, error)
	FindUserByEmail(context.Context, string) (entity.User, error)
}
