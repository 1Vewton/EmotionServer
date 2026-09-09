package utilapi

import (
	"net/http"

	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/gin-gonic/gin"
)

// CheckHealth checks the health status
// @Summary Checks health
// @Schemes
// @Description check health
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} api.Response
// @Router /v1/utils/health [get]
func CheckHealth(c *gin.Context) {
	response.NewResponse(
		c,
		http.StatusOK,
		true,
		nil,
		nil,
	)
}
