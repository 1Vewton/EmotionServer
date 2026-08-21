package main

import (
	"io"
	"net/http"
	"os"

	"github.com/1Vewton/EmotionServer/api/utilapi"
	"github.com/gin-gonic/gin"
)

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
	utilRouter := router.Group("/utils")
	utilRouter.GET(
		"/health",
		utilapi.CheckHealth,
	)
	// Define server
	s := &http.Server{
		Handler: router,
	}
	s.ListenAndServe()
}
