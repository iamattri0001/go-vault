package auth

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/pkg/jwt"
	"go-vault/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service *service.Service
}

func NewLoginController(service *service.Service) *LoginController {
	return &LoginController{
		service: service,
	}
}

func (c *LoginController) GetResponse(ctx *gin.Context) {
	var request service.LoginUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	user, err := c.service.LoginUser(&request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	token, err := jwt.GenerateToken(user.Username, 24)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}
	ctx.SetCookie("token", token, 3600, "/", "", false, true)
	controller.SendResponse(ctx, true, "User logged in successfully", map[string]any{"user": user}, nil)
}
