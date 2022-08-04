package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/dto"
)

type CommentInteractor interface {
	CreateComment(context.Context, dto.CreateCommentReq) error
	FindCommentsByPostID(context.Context, dto.FindCommentReq) (dto.FindCommentRes, error)
	DeleteUserComments(context.Context, dto.DeleteUserCommentsReq) error
}
