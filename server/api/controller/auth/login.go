package auth

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/pkg/hash"
	"go-vault/pkg/jwt"
	"go-vault/service"
	"log"

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

	if err := hash.CheckHash(request.Password, user.PasswordHash); err != nil {
		log.Printf("Login failed for user %s: %v", request.Username, err)
		controller.SendResponse(ctx, false, "Invalid username or password", nil, customerrors.ErrInvalidCredentials)
		return
	}

	token, err := jwt.GenerateToken(user, 24)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}
	ctx.SetCookie("token", token, 3600, "/", "", false, true)
	controller.SendResponse(ctx, true, "User logged in successfully", map[string]any{"user": user}, nil)
}
