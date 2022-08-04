package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mohammadmokh/writino/dto"
)

func ValidateRegisterUser(req dto.RegisterReq) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 0)))
}

func ValidateUpdateUser(req dto.UpdateUserReq) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Bio, validation.Length(0, 500)),
		validation.Field(&req.DisplayName, validation.Length(0, 256)),
		validation.Field(&req.ProfilePic, is.URL),
	)
}

func ValidateUpdatePassword(req dto.UpdatePasswordReq) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Old, validation.Required),
		validation.Field(&req.New, validation.Required, validation.Length(8, 0)))
}
