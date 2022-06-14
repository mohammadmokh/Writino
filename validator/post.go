package validator

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/dto"
)

func ValidateCreatePost(req dto.CreatePostReq) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 500)),
		validation.Field(&req.Tags, validation.Length(0, 3), validation.Each(validation.Length(3, 20))),
		validation.Field(&req.Content, validation.Required),
	)
}

func ValidateUpdatePost(req dto.UpdatePostReq) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Length(3, 500)),
		validation.Field(&req.Tags, validation.Length(0, 3), validation.Each(validation.Length(3, 20))),
	)
}

func ValidateDeletePost(ctx context.Context, req dto.DeletePostReq, store contract.PostStore) error {
	return validation.Validate(req, validation.By(deletePost(ctx, store)))
}
