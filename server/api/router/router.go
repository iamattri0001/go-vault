package router

import (
	"go-vault/service"

	"go-vault/api/controller"
	"go-vault/api/controller/auth"
	"go-vault/api/controller/vault"
	"go-vault/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service *service.Service) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	addUserRoutes(r, service)
	addVaultRoutes(r, service)
}

func addUserRoutes(r *gin.Engine, service *service.Service) {
	authGrp := r.Group("/api/v1/auth")
	{
		authGrp.POST("/register", auth.NewRegisterController(service).GetResponse)
		authGrp.POST("/login", auth.NewLoginController(service).GetResponse)
		authGrp.GET("/logout", func(ctx *gin.Context) {
			ctx.SetCookie("token", "", -1, "/", "", false, true)
			controller.SendResponse(ctx, true, "User logged out successfully", nil, nil)
		})
	}
}

func addVaultRoutes(r *gin.Engine, service *service.Service) {
	vaultGrp := r.Group("/api/v1/vault").Use(middleware.AuthMiddleware())
	{
		vaultGrp.POST("/create", vault.NewCreateVaultController(service).GetResponse)
		vaultGrp.GET("/list", vault.NewListVaultsController(service).GetResponse)
		vaultGrp.POST("/:id", vault.NewUpdateVaultController(service).GetResponse)
		// vaultGrp.GET("/:id", auth.NewGetVaultController(service).GetResponse)
		// vaultGrp.DELETE("/:id", auth.NewDeleteVaultController(service).GetResponse)
	}
}
