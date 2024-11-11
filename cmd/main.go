package main

import (
	"fmt"
	"go-simple-project/internal/common/dependencies"
	"go-simple-project/internal/entrypoints"
	"go-simple-project/internal/infrastructure"
	"go-simple-project/internal/services"

	_ "go-simple-project/docs"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func initializeDependencies() *dependencies.Dependency {
	logger.Info("Starting initialize dependencies...")
	deps, err := dependencies.InitializeDependencies()
	if err != nil {
		logger.Fatalf("Failed to initialize dependencies: %v", err)
	}
	logger.Info("End initialize dependencies.")
	return deps
}

func initializeSearchValueService(
	deps *dependencies.Dependency,
	repository infrastructure.IFileNumberRepository,
) *services.SearchValueService {
	service, err := services.NewSearchValueService(repository, deps)
	if err != nil {
		logger.Fatalf("Failed to initialize 'SearchValueService' service: %v", err)
	}
	logger.Info("Data has been successfully loaded.")
	return service
}

func initializeHttpServer(
	svc *services.SearchValueService,
	deps *dependencies.Dependency,
) {
	logger.Info("Starting HTTP server...")
	router := gin.Default()
	entrypoints.SetupRoutes(router, svc)
	router.Run(fmt.Sprintf(":%v", deps.Config.HttpServer.Port))
}

func main() {
	logger.Info("Starting application...")
	deps := initializeDependencies()
	repository := infrastructure.NewFileNumberRepository(deps)
	svc := initializeSearchValueService(deps, repository)
	initializeHttpServer(svc, deps)
}
