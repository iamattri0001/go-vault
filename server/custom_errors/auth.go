package customerrors

import "errors"

var (
	ErrInvalidUsernameFormat = errors.New("username must be between 3 and 15 characters")
	ErrUsernameTaken         = errors.New("username is already taken")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUnauthorized          = errors.New("unauthorized")
)
