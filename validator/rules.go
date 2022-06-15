package validator

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/dto"
	errs "gitlab.com/gocastsian/writino/errors"
)

func deletePost(ctx context.Context, store contract.PostStore) validation.RuleFunc {

	return func(value interface{}) error {

		req, _ := value.(dto.DeletePostReq)
		post, err := store.FindPostByID(ctx, req.ID)

		if err == errs.ErrNotFound {
			return nil
		}

		if post.AuthorID != req.UserID {
			return errors.New("can't delete this post")
		}
		return nil
	}
}
