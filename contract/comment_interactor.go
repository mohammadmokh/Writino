package contract

import (
	"context"

	"github.com/mohammadmokh/writino/dto"
)

type CommentInteractor interface {
	CreateComment(context.Context, dto.CreateCommentReq) error
	FindCommentsByPostID(context.Context, dto.FindCommentReq) (dto.FindCommentRes, error)
	DeleteUserComments(context.Context, dto.DeleteUserCommentsReq) error
}
