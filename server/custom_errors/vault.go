package customerrors

import "errors"

var (
	ErrInvalidVaultTitleFormat       = errors.New("vault title must be between 3 and 30 characters")
	ErrInvalidVaultDescriptionFormat = errors.New("vault description must be 100 characters or less")
	ErrVaultNotFound                 = errors.New("vault not found")
	ErrVaultAlreadyExists            = errors.New("vault already exists")
)
