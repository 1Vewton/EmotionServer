package dataapi

import (
	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/gin-gonic/gin"
)

// AddAgentProfile adds agent profile.
// @Summary Adds profile for agent
// @Schemes
// @Description Adds profile for agent
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} api.Response
// @Router /v1/data/addAgentProfile [post]
// @Param req body NewAgentProfileQuery true "Query parameters"
func AddAgentProfile(
	c *gin.Context,
) {
	var query NewAgentProfileQuery
	err := c.ShouldBindJSON(&query)
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
	personality := ocean.NewPersonality(
		query.Openness,
		query.Conscientiousness,
		query.Extraversion,
		query.Agreeableness,
		query.Neuroticism,
	)
	err = manager.MainAgentProfileManager.AddNewAgentProfile(
		c,
		personality,
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
	response.NewResponse(
		c,
		201,
		true,
		nil,
		nil,
	)
}
