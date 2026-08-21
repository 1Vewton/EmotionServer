package utilapi

import (
	"net/http"

	"github.com/1Vewton/EmotionServer/api"
	"github.com/gin-gonic/gin"
)

// CheckHealth checks the health status
func CheckHealth(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		api.NewResponse(
			true,
			nil,
			nil,
		),
	)
}
