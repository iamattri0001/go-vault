package auth

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	service *service.Service
}

func NewRegisterController(service *service.Service) *RegisterController {
	return &RegisterController{
		service: service,
	}
}

func (c *RegisterController) GetResponse(ctx *gin.Context) {
	var request service.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	user, err := c.service.CreateUser(&request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	controller.SendResponse(ctx, true, "User registered successfully", map[string]any{"user": user}, nil)
}
