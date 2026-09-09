package api

import (
	"github.com/1Vewton/EmotionServer/api/dataapi"
	"github.com/1Vewton/EmotionServer/api/utilapi"
	"github.com/1Vewton/EmotionServer/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter sets up router
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	apiV1 := router.Group("/v1")
	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	utilRouter := apiV1.Group("/utils")
	utilRouter.GET(
		"/health",
		utilapi.CheckHealth,
	)
	dataRouter := apiV1.Group("/data")
	dataRouter.POST(
		"/addAgentProfile",
		dataapi.AddAgentProfile,
	)
	return router
}
