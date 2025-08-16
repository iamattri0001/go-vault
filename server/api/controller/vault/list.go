package vault

import (
	"go-vault/api/controller"
	"go-vault/service"
	"go-vault/utils"

	"github.com/gin-gonic/gin"
)

type ListVaultsController struct {
	service *service.Service
}

func NewListVaultsController(service *service.Service) *ListVaultsController {
	return &ListVaultsController{
		service: service,
	}
}

func (c *ListVaultsController) GetResponse(ctx *gin.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}
	vaults, err := c.service.ListVaults(userID)
	if err != nil {
		controller.SendResponse(ctx, false, "", nil, err)
		return
	}
	controller.SendResponse(ctx, true, "Vaults retrieved successfully", map[string]any{"vaults": vaults}, nil)
}
