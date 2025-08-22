package customerrors

import "errors"

var (
	ErrPasswordNotFound           = errors.New("password not found")
	ErrPasswordAlreadyExists      = errors.New("password already exists")
	ErrInvalidPasswordTitleFormat = errors.New("invalid password title format")
)
