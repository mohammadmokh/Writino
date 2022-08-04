package comment

import (
	"context"

	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/dto"
	"github.com/mohammadmokh/writino/entity"
)

type CommentInteractor struct {
	store contract.CommentStore
}

func New(store contract.CommentStore) contract.CommentInteractor {
	return CommentInteractor{
		store: store,
	}
}

func (i CommentInteractor) CreateComment(ctx context.Context, req dto.CreateCommentReq) error {

	err := i.store.CreateComment(ctx, entity.Comment{
		UserID: req.UserID,
		Text:   req.Comment,
	}, req.PostID)

	return err
}

func (i CommentInteractor) FindCommentsByPostID(ctx context.Context, req dto.FindCommentReq) (
	dto.FindCommentRes, error) {

	findRes, err := i.store.FindCommentsByPostID(ctx, contract.FindCommentfilters{
		PostID: req.PostID,
		Limit:  req.Limit,
		Page:   req.Page,
	})
	if err != nil {
		return dto.FindCommentRes{}, err
	}

	res := dto.FindCommentRes{TotalCount: findRes.TotalCount}

	for i := 0; len(findRes.Comments) > i; i++ {
		comment := dto.Comment{
			Comment: findRes.Comments[i].Text,
			User:    findRes.Comments[i].UserID,
		}
		res.Comments = append(res.Comments, comment)
	}

	return res, nil
}

func (i CommentInteractor) DeleteUserComments(ctx context.Context, req dto.DeleteUserCommentsReq) error {

	return i.store.DeleteCommentsByUserID(ctx, req.UserID)
}
