package password

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordController struct {
	service *service.Service
}

func NewUpdatePasswordController(service *service.Service) *UpdatePasswordController {
	return &UpdatePasswordController{
		service: service,
	}
}

func (c *UpdatePasswordController) GetResponse(ctx *gin.Context) {
	var request *service.UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	password, err := c.service.UpdatePassword(userID, request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	controller.SendResponse(ctx, true, "Password created successfully", map[string]any{"password": password}, nil)
}
