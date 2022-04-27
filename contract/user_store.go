package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type UserStore interface {
	CreateUser(context.Context, entity.User) (entity.User, error)
	FindUser(context.Context, string) (entity.User, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, string) error
}
