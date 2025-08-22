package vault

import (
	"go-vault/api/controller"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteVaultController struct {
	service *service.Service
}

func NewDeleteVaultController(service *service.Service) *DeleteVaultController {
	return &DeleteVaultController{
		service: service,
	}
}

func (c *DeleteVaultController) GetResponse(ctx *gin.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "Failed to get user ID", nil, err)
		return
	}
	vaultID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		controller.SendResponse(ctx, false, "Invalid vault ID", nil, err)
		return
	}
	err = c.service.DeleteVaultAndPasswordByVaultID(userID, vaultID)
	if err != nil {
		controller.SendResponse(ctx, false, "Failed to delete vault", nil, err)
		return
	}
	controller.SendResponse(ctx, true, "Vault deleted successfully", nil, nil)
}
