package main

import (
	"io"
	"os"

	"github.com/1Vewton/EmotionServer/api/utilapi"
	"github.com/1Vewton/EmotionServer/docs"
	"github.com/1Vewton/EmotionServer/pkg/settings"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Emotion Simulator
// @version 1.0
// @description A service to simulate emotion for agent
// @termsOfService http://swagger.io/terms/

// @contact.name 1Vewton
// @contact.url https://github.com/1Vewton
// @contact.email zhanyunze0601@gmai.com

// @license.name MIT
func main() {
	gin.DisableConsoleColor()
	file, err := os.Create("server.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(
		file,
		os.Stdout,
	)
	// Define router
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	utilRouter := router.Group("/utils")
	utilRouter.GET(
		"/health",
		utilapi.CheckHealth,
	)
	// Define server
	router.Run(
		settings.Settings.GetServerURL(),
	)
}
