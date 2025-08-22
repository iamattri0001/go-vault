package password

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
)

type CreatePasswordController struct {
	service *service.Service
}

func NewCreatePasswordController(service *service.Service) *CreatePasswordController {
	return &CreatePasswordController{
		service: service,
	}
}

func (c *CreatePasswordController) GetResponse(ctx *gin.Context) {
	var request *service.CreatePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	password, err := c.service.CreatePasswordWithVaultID(userID, request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	controller.SendResponse(ctx, true, "Password created successfully", map[string]any{"password": password}, nil)
}
