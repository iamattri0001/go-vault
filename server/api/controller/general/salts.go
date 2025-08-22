package general

import (
	"go-vault/api/controller"
	"go-vault/service"

	"github.com/gin-gonic/gin"
)

type GetSaltsController struct {
	svc *service.Service
}

func NewGetSaltsController(svc *service.Service) *GetSaltsController {
	return &GetSaltsController{
		svc: svc,
	}
}

func (c *GetSaltsController) GetResponse(ctx *gin.Context) {
	username := ctx.Param("username")
	salts, err := c.svc.GetSalts(username)
	if err != nil {
		controller.SendResponse(ctx, false, "Failed to get salts", nil, err)
		return
	}
	controller.SendResponse(ctx, true, "Salts retrieved successfully", salts, nil)
}
