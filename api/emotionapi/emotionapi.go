package emotionapi

import (
	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/gin-gonic/gin"
)

// RegisterEmotion registers emotion context
// @Summary Registers new emotion while returning a context id in the result
// @Schemes
// @Description Registers new emotion while returning a context id in the result
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/data/addAgentProfile [post]
// @Param req body NewContextQuery true "Query parameters"
func RegisterEmotion(
	c *gin.Context,
) {
	var query NewContextQuery
	err := c.ShouldBindJSON(
		&query,
	)
	if err != nil {
		response.NewResponse(
			c,
			400,
			false,
			nil,
			err,
		)
		return
	}
}
