package customerrors

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrBadRequest         = errors.New("bad request")
)

var ErrMap = map[error]int{
	ErrSomethingWentWrong:    500,
	ErrBadRequest:            400,
	ErrInvalidPasswordFormat: 400,
	ErrInvalidUsernameFormat: 400,
	ErrUsernameTaken:         400,
	ErrUserNotFound:          404,
	ErrInvalidCredentials:    401,
}

func GetCode(err error) int {
	if code, ok := ErrMap[err]; ok {
		return code
	}
	return 500
}
