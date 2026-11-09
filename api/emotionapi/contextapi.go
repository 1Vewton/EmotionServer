package emotionapi

import (
	"errors"

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
		return
	}
	// Checks if profile exists and whether it exists in redis, if not exists while existing in redis, delete it
	profileExists, err := manager.MainAgentProfileManager.Exists(
		c.Request.Context(),
		query.AgentKey,
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
	if !profileExists {
		hasProfileInRedis, err := manager.MainContextManager.HasProfile(
			c.Request.Context(),
			query.AgentKey,
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
		if hasProfileInRedis {
			err = manager.MainContextManager.DeleteProfileData(
				c.Request.Context(),
				query.AgentKey,
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
		}
		response.NewResponse(
			c,
			404,
			false,
			nil,
			errors.New(
				"such profile does not exist",
			),
		)
		return
	}
	err = manager.MainContextManager.CleanAllData(
		c.Request.Context(),
		query.AgentKey,
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
	contexts, err := manager.MainContextManager.GetContextsInProfile(
		c.Request.Context(),
		query.AgentKey,
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
	resp := &GetContextsOfAProfileResponse{
		Contexts: contexts,
	}
	response.NewResponse(
		c,
		200,
		true,
		resp,
		nil,
	)
}
