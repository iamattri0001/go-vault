package router

import (
	"go-vault/service"

	"go-vault/api/controller"
	"go-vault/api/controller/auth"
	"go-vault/api/controller/general"
	"go-vault/api/controller/password"
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
	addPasswordRoutes(r, service)
}

func addUserRoutes(r *gin.Engine, service *service.Service) {
	// general routes which don't require authentication
	r.GET("/api/v1/salts/:username", general.NewGetSaltsController(service).GetResponse)

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
		vaultGrp.PUT("/:id", vault.NewUpdateVaultController(service).GetResponse)
		vaultGrp.GET("/:id", vault.NewGetVaultController(service).GetResponse)
		vaultGrp.DELETE("/:id", vault.NewDeleteVaultController(service).GetResponse)
	}
}

func addPasswordRoutes(r *gin.Engine, service *service.Service) {
	passwordGrp := r.Group("/api/v1/password").Use(middleware.AuthMiddleware())
	{
		passwordGrp.POST("/create", password.NewCreatePasswordController(service).GetResponse)
		// passwordGrp.GET("/list", password.NewListPasswordsController(service).GetResponse)
		// passwordGrp.PUT("/:id", password.NewUpdatePasswordController(service).GetResponse)
		// passwordGrp.DELETE("/:id", password.NewDeletePasswordController(service).GetResponse)
	}
}
