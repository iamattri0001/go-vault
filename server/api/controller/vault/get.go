package vault

import (
	"go-vault/api/controller"
	customerrors "go-vault/custom_errors"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetVaultController struct {
	service *service.Service
}

func NewGetVaultController(service *service.Service) *GetVaultController {
	return &GetVaultController{
		service: service,
	}
}

func (c *GetVaultController) GetResponse(ctx *gin.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}

	vaultID, err := uuid.Parse(ctx.Param("id"))
	if err != nil || vaultID == uuid.Nil {
		controller.SendResponse(ctx, false, "", nil, customerrors.ErrBadRequest)
		return
	}
	vaultPasswords, err := c.service.GetVaultPasswords(userID, vaultID)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}
	controller.SendResponse(ctx, true, "Vault passwords retrieved successfully", map[string]any{"passwords": vaultPasswords}, nil)
}
