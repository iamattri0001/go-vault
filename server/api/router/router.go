package router

import (
	"go-vault/service"

	"go-vault/api/controller/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service *service.Service) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	addUserRoutes(r, service)

}

func addUserRoutes(r *gin.Engine, service *service.Service) {
	authGrp := r.Group("/api/v1/auth")
	{
		authGrp.POST("/register", auth.NewRegisterController(service).GetResponse)
	}
}
