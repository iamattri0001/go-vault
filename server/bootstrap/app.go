package bootstrap

import (
	"fmt"
	"go-vault/api/router"
	"go-vault/config"
	"go-vault/database/mongodb"
	"go-vault/injector"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config    *config.Config
	GinEngine *gin.Engine
	MongoDB   *mongodb.MongoDBConnection
	Injector  *injector.DependencyInjector
}

func NewApp() (*App, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	mongoDbConnection, err := mongodb.NewMongoDbConnection(&config.MongoConfig)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	di := injector.NewDependencyInjector(injector.WithAppConfig(config), injector.WithMongoDB(mongoDbConnection))

	return &App{
		Config:    config,
		GinEngine: gin.Default(),
		MongoDB:   mongoDbConnection,
		Injector:  di,
	}, nil
}

func (a *App) Run() {
	defer a.Close()
	port := a.Config.ServiceConfig.Port
	address := fmt.Sprintf(":%d", port)
	router.SetupRoutes(a.GinEngine, a.Injector)
	if err := a.GinEngine.Run(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func (a *App) Close() {
	if err := a.MongoDB.Close(); err != nil {
		log.Fatalf("Failed to close MongoDB connection: %v", err)
	}
}
