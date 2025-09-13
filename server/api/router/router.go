package router

import (
	"go-vault/injector"
	"go-vault/service"
	"time"

	"go-vault/api/controller"
	"go-vault/api/controller/auth"
	"go-vault/api/controller/general"
	"go-vault/api/controller/password"
	"go-vault/api/controller/vault"
	"go-vault/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, di *injector.DependencyInjector) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     di.AppConfig.ServiceConfig.Clients,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	addUserRoutes(r, di.Service)
	addVaultRoutes(r, di.Service)
	addPasswordRoutes(r, di.Service)
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
		vaultGrp.POST("/", vault.NewCreateVaultController(service).GetResponse)
		vaultGrp.GET("/list", vault.NewListVaultsController(service).GetResponse)
		vaultGrp.PUT("/", vault.NewUpdateVaultController(service).GetResponse)
		vaultGrp.GET("/:id", vault.NewGetVaultController(service).GetResponse)
		vaultGrp.DELETE("/:id", vault.NewDeleteVaultController(service).GetResponse)
	}
}

func addPasswordRoutes(r *gin.Engine, service *service.Service) {
	passwordGrp := r.Group("/api/v1/password").Use(middleware.AuthMiddleware())
	{
		passwordGrp.POST("/", password.NewCreatePasswordController(service).GetResponse)
		passwordGrp.PUT("/", password.NewUpdatePasswordController(service).GetResponse)
		// passwordGrp.DELETE("/:id", password.NewDeletePasswordController(service).GetResponse)
	}
}
