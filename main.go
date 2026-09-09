//go:build !gRPC

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/1Vewton/EmotionServer/api"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/logger"
	"github.com/1Vewton/EmotionServer/pkg/settings"
	"github.com/gin-gonic/gin"
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
	db, err := database.Connect(
		settings.Settings.GetDatabaseURL(),
		settings.Settings.GetDatabaseType(),
		&profile.AgentProfile{},
	)
	if err != nil {
		panic(err)
	}
	database.DB = db
	database.RedisClient = database.InitRedisClient(
		settings.Settings.GetDatabaseURL(),
		settings.Settings.GetRedisPassword(),
		settings.Settings.GetRedisDialTimeout(),
		settings.Settings.GetRedisReadTimeout(),
		settings.Settings.GetRedisWriteTimeout(),
		settings.Settings.GetRedisMaxRetries(),
		settings.Settings.GetRedisMaxRetryBackOff(),
		settings.Settings.GetRedisMinRetryBackOff(),
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
	router := api.SetUpRouter()
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
	err = srv.Shutdown(ctx)
	if err != nil {
		logger.SysLogger.Error(err.Error())
	}
	logger.SysLogger.Info("Start closing database connection")
	err = database.Close(database.DB)
	if err != nil {
		logger.SysLogger.Error(err.Error())
	}
	logger.SysLogger.Info("Start closing redis connection")
	err = database.CloseRedis(
		database.RedisClient,
	)
	if err != nil {
		logger.SysLogger.Error(err.Error())
	}
}
