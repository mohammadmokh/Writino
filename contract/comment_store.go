package contract

import (
	"context"

	"github.com/mohammadmokh/writino/entity"
)

type FindCommentfilters struct {
	PostID string
	Limit  int
	Page   int
}

type FindCommentRes struct {
	TotalCount int
	Comments   []entity.Comment
}

type CommentStore interface {
	CreateComment(context.Context, entity.Comment, string) error
	FindCommentsByPostID(context.Context, FindCommentfilters) (FindCommentRes, error)
	DeleteCommentsByUserID(context.Context, string) error
}
