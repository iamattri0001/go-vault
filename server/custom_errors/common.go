package customerrors

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrBadRequest         = errors.New("bad request")
)

var (
	ErrInvalidTitleFormat       = errors.New("title must be between 3 and 30 characters")
	ErrInvalidDescriptionFormat = errors.New("description must be 100 characters or less")
)
