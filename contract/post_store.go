package contract

import (
	"context"

	"github.com/mohammadmokh/writino/entity"
)

type SearchPostFilters struct {
	Query string
	Limit int
	Page  int
}

type SearchPostRes struct {
	Count int
	Posts []entity.Post
}

type PostStore interface {
	CreatePost(context.Context, entity.Post) (entity.Post, error)
	FindPostByID(context.Context, string) (entity.Post, error)
	FindPostsByUserID(context.Context, SearchPostFilters) (SearchPostRes, error)
	UpdatePost(context.Context, entity.Post) error
	DeletePost(context.Context, string) error
	SearchPost(ctx context.Context, filters SearchPostFilters) (SearchPostRes, error)
	FindAll(context.Context, SearchPostFilters) (SearchPostRes, error)
	LikePost(ctx context.Context, postID string, userID string) error
	DeletePostsByUserID(context.Context, string) error
}
