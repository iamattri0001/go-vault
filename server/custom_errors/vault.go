package customerrors

import "errors"

var (
	ErrVaultNotFound           = errors.New("vault not found")
	ErrVaultAlreadyExists      = errors.New("vault already exists")
	ErrInvalidVaultTitleFormat = errors.New("invalid vault title format")
)
