package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/1Vewton/EmotionServer/api/utilapi"
	"github.com/1Vewton/EmotionServer/docs"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/logger"
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
	ctx := context.Background()
	database.Connect(
		settings.Settings.GetDatabaseURL(),
		settings.Settings.GetDatabaseType(),
		&profile.AgentProfile{},
	)
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
	address := settings.Settings.GetServerURL()
	logger.SysLogger.Info(
		fmt.Sprintf(
			"Running on %s",
			address,
		),
	)
	srv := &http.Server{
		Addr:    address,
		Handler: router.Handler(),
	}
	runServer := func() {
		err := srv.ListenAndServe()
		if err != nil {
			logger.SysLogger.Error(err.Error())
			os.Exit(2)
		}
	}
	go runServer()
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	logger.SysLogger.Info("Start closing program")
	srv.Shutdown(ctx)
	logger.SysLogger.Info("Start cleaning database")
	database.Close()
}
