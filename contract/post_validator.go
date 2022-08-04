package contract

import (
	"github.com/mohammadmokh/writino/dto"
)

type (
	ValidateCreatePost func(req dto.CreatePostReq) error
	ValidateUpdatePost func(req dto.UpdatePostReq) error
)
