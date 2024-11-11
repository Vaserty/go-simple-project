package entrypoints

import (
	"context"
	"go-simple-project/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, searchValueService *services.SearchValueService) {
	router.Use(func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "SearchValueService", searchValueService)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	numbersGroup := router.Group("/numbers")
	numbersGroup.GET("/:value", SearchIndexEndpoint)
}
