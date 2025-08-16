package vault

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
)

type CreateVaultController struct {
	service *service.Service
}

func NewCreateVaultController(service *service.Service) *CreateVaultController {
	return &CreateVaultController{
		service: service,
	}
}

func (c *CreateVaultController) GetResponse(ctx *gin.Context) {
	var request service.CreateVaultRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	vault, err := c.service.CreateVault(userID, &request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	controller.SendResponse(ctx, true, "Vault created successfully", map[string]any{"vault": vault}, nil)
}
