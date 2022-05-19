package errors

import "errors"

var (
	ErrNotFound           = errors.New("record not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateRecord    = errors.New("duplicate value")
)
