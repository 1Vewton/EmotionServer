package emotionapi

import (
	"fmt"
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
	exists, err := manager.MainAgentProfileManager.Exists(
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
	if !exists {
		response.NewResponse(
			c,
			404,
			false,
			nil,
			fmt.Errorf(
				"%s api key does not exists",
				query.APIKey,
			),
		)
	}
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

// GetCurrentEmotion gets current emotion
// @Summary Getting through api key and context key
// @Schemes
// @Description Fetching current emotion(the last time updated) through api key and context key, returning new context key and initial emotion if context invalid, returning emotion if context is valid
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/emotion/getCurrentEmotion [post]
// @Param req body GetCurrentEmotionQuery true "Query parameters"
func GetCurrentEmotion(
	c *gin.Context,
) {
	var query GetCurrentEmotionQuery
	err := c.ShouldBindJSON(&query)
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
	// Checks agent existence
	agentExists, err := manager.MainAgentProfileManager.Exists(
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
		return
	}
	if !agentExists {
		response.NewResponse(
			c,
			404,
			false,
			nil,
			fmt.Errorf(
				"agent with api key %s does not exists",
				query.APIKey,
			),
		)
		return
	}
	// Checks context existence
	contextExistsInProfile, err := manager.MainContextManager.ContextExistsWithAgentProfile(
		c.Request.Context(),
		query.APIKey,
		query.ContextKey,
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
	contextExistsInDatabase, err := manager.MainContextManager.HasContext(
		c.Request.Context(),
		query.ContextKey,
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
	contextExists := contextExistsInDatabase && contextExistsInProfile
	agentProfile, err := manager.MainAgentProfileManager.SearchAgentProfile(
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
		return
	}
	if !contextExists {
		contextID := uuid.NewString()
		contextKey, err := manager.MainContextManager.NewContext(
			c.Request.Context(),
			query.APIKey,
			contextID,
			agentProfile.ContextExistsTime,
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
		resp := &GetCurrentEmotionResponse{
			CurrentEmotion: agentProfile.InitialEmotion,
			NewContextKey:  &contextKey,
		}
		response.NewResponse(
			c,
			201,
			true,
			resp,
			nil,
		)
		return
	}
	contextData, err := manager.MainContextManager.GetDataFromContext(
		c.Request.Context(),
		query.ContextKey,
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
	resp := &GetCurrentEmotionResponse{
		CurrentEmotion: contextData.CurrentEmotion,
		NewContextKey:  nil,
	}
	response.NewResponse(
		c,
		200,
		true,
		resp,
		nil,
	)
}
