package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
)

type PostStore interface {
	CreatePost(context.Context, entity.Post) (entity.Post, error)
	FindPostByID(context.Context, string) (entity.Post, error)
	FindPostsByUserID(context.Context, string) ([]entity.Post, error)
	UpdatePost(context.Context, entity.Post) error
	DeletePost(context.Context, string) error
}
