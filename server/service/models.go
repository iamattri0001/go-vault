package service

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateVaultRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type UpdateVaultRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}
