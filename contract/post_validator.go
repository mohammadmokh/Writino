package contract

import (
	"context"

	"gitlab.com/gocastsian/writino/dto"
)

type (
	ValidateCreatePost func(req dto.CreatePostReq) error
	ValidateUpdatePost func(req dto.UpdatePostReq) error
	ValidateDeletePost func(ctx context.Context, req dto.DeletePostReq, store PostStore) error
)
