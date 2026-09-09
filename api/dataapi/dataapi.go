package dataapi

import (
	"github.com/1Vewton/EmotionServer/api"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/pkg/database"
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
		api.NewResponse(
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
	err = database.AddNewAgentProfile(
		c,
		personality,
		query.APIKey,
	)
	if err != nil {
		api.NewResponse(
			c,
			500,
			false,
			nil,
			err,
		)
		return
	}
	api.NewResponse(
		c,
		201,
		true,
		nil,
		nil,
	)
}
