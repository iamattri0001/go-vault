package customerrors

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrBadRequest         = errors.New("bad request")
)
