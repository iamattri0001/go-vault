package customerrors

var ErrMap = map[error]int{
	ErrSomethingWentWrong:            500,
	ErrBadRequest:                    400,
	ErrInvalidPasswordFormat:         400,
	ErrInvalidUsernameFormat:         400,
	ErrUsernameTaken:                 400,
	ErrUserNotFound:                  404,
	ErrInvalidCredentials:            401,
	ErrInvalidVaultTitleFormat:       400,
	ErrInvalidVaultDescriptionFormat: 400,
	ErrUnauthorized:                  401,
	ErrVaultNotFound:                 404,
}

func GetCode(err error) int {
	if code, ok := ErrMap[err]; ok {
		return code
	}
	return 500
}
