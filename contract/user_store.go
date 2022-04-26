package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type UserStore interface {
	Create(context.Context, entity.User) (entity.User, error)
	FindbyID(context.Context, string) (entity.User, error)
	Update(context.Context, entity.User) error
	Delete(context.Context, string) error
}
