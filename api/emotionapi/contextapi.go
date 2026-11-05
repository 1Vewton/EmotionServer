package emotionapi

import (
	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/gin-gonic/gin"
)

// GetContextsOfAProfile gets all the contexts created for a profile
// @Summary Getting contexts through profile api key
// @Schemes
// @Description Getting contexts through profile api key
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/emotion/getContextsOfAProfile [post]
// @Param req body GetCurrentEmotionQuery true "Query parameters"
func GetContextsOfAProfile(
	c *gin.Context,
) {
	var query GetContextsOfAProfileQuery
	err := c.BindJSON(&query)
	if err != nil {
		response.NewResponse(
			c,
			500,
			false,
			nil,
			err,
		)
	}
	err = manager.MainContextManager.CleanAllData(
		c.Request.Context(),
		query.APIKey,
	)
	if err != nil {
		response.NewResponse(
			c,
			500,
			false,
			nil,
			err,
		)
	}
}
