package emotionapi

import (
	"time"

	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterEmotion registers emotion context
// @Summary Registers new emotion while returning a context id in the result
// @Schemes
// @Description Registers new emotion while returning a context id in the result
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/emotion/registerEmotion [post]
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
	contextID := uuid.NewString()
	time := 24 * time.Hour
	contextKey, err := manager.MainContextManager.NewContext(
		c.Request.Context(),
		query.APIKey,
		contextID,
		&time,
	)
	if err != nil {
		response.NewResponse(
			c,
			500,
			false,
			nil,
			err,
		)
		return
	}
	err = manager.MainContextManager.AddNewContextForProfile(
		c.Request.Context(),
		query.APIKey,
		contextKey,
	)
	if err != nil {
		response.NewResponse(
			c,
			500,
			false,
			nil,
			err,
		)
		return
	}
	resp := &WithContextResponse{
		ContextKey: contextKey,
	}
	response.NewResponse(
		c,
		201,
		true,
		resp,
		nil,
	)
}
