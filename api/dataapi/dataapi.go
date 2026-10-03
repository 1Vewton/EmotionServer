package dataapi

import (
	"fmt"
	"time"

	"github.com/1Vewton/EmotionServer/api/response"
	"github.com/1Vewton/EmotionServer/internal/manager"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/settings"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddAgentProfile adds agent profile.
// @Summary Adds profile for agent
// @Schemes
// @Description Adds profile for agent
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
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
	var defaultTime time.Duration
	if query.ContextLifeTimeInDays == nil {
		defaultTime = settings.Settings.GetContextLastTimeInDays()
	} else {
		defaultTime = time.Duration(*query.ContextLifeTimeInDays*24) * time.Hour
	}
	apiKey := uuid.NewString()
	result := &NewAgentProfileResponse{
		APIKey: apiKey,
	}
	newAgentProfile, err := profile.NewAgentProfile(
		apiKey,
		personality,
		defaultTime,
	)
	newAgentProfile = newAgentProfile.SetIsContextInfinite(
		query.IsContextInfinite,
	)
	err = manager.MainAgentProfileManager.AddNewAgentProfile(
		c.Request.Context(),
		newAgentProfile,
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
		result,
		nil,
	)
}

// GetAgentProfile searches certain agent profile through api key
// @Summary Searches profile for agent
// @Schemes
// @Description Searches profile for agent
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /v1/data/getAgentProfile [post]
// @Param req body GetAgentProfileQuery true "Query parameters"
func GetAgentProfile(
	c *gin.Context,
) {
	var query GetAgentProfileQuery
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
		return
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
	searchResult, err := manager.MainAgentProfileManager.SearchAgentProfile(
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
	response.NewResponse(
		c,
		200,
		true,
		searchResult,
		nil,
	)
}
