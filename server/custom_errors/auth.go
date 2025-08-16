package customerrors

import "errors"

var (
	ErrInvalidUsernameFormat = errors.New("username must be between 3 and 15 characters")
	ErrInvalidPasswordFormat = errors.New("password must be between 8 and 30 characters")
	ErrUsernameTaken         = errors.New("username is already taken")
)
