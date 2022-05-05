package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type CommentStore interface {
	CreateComment(context.Context, entity.Comment, string) error
	FindByPostID(context.Context, string) ([]entity.Comment, error)
}
