package contract

import (
	"gitlab.com/gocastsian/writino/dto"
)

type (
	ValidateCreatePost func(req dto.CreatePostReq) error
	ValidateUpdatePost func(req dto.UpdatePostReq) error
)
