package vault

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateVaultController struct {
	service *service.Service
}

func NewUpdateVaultController(service *service.Service) *UpdateVaultController {
	return &UpdateVaultController{
		service: service,
	}
}

func (c *UpdateVaultController) GetResponse(ctx *gin.Context) {
	var request service.UpdateVaultRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.SendResponse(ctx, false, "Invalid request format", nil, customerrors.ErrBadRequest)
		return
	}

	vaultID, err := uuid.Parse(ctx.Param("id"))
	if err != nil || vaultID == uuid.Nil {
		controller.SendResponse(ctx, false, "Vault ID is required", nil, customerrors.ErrBadRequest)
		return
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	vault, err := c.service.UpdateVault(userID, vaultID, &request)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	controller.SendResponse(ctx, true, "Vault updated successfully", map[string]any{"vault": vault}, nil)
}
