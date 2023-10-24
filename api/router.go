package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	handler "gitlab.com/iman_api/api/handlers"
	"gitlab.com/iman_api/pkg/logger"
)

type RoutetOptions struct {
	Log logger.Logger
}

// @Description Created by Otajonov Quvonchbek
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RoutetOptions) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running...",
		})
	})

	handler := handler.New(&handler.HandlerOptions{
		Log: option.Log,
	})

	api := router.Group("/v1")

	api.GET("/days", handler.AuthMiddleWare, handler.Days)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
