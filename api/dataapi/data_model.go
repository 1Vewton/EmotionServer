package dataapi

import (
	"time"

	"github.com/1Vewton/EmotionServer/internal/emotion"
	"github.com/1Vewton/EmotionServer/internal/ocean"
)

// NewAgentProfileQuery defines the data structure for creating new agent profile
type NewAgentProfileQuery struct {
	Openness              float64
	Conscientiousness     float64
	Extraversion          float64
	Agreeableness         float64
	Neuroticism           float64
	IsContextInfinite     bool
	ContextLifeTimeInDays *int
}

// NewAgentProfileResponse defines the response for creating new agent profile response
type NewAgentProfileResponse struct {
	APIKey string
}

// GetAgentProfileQuery defines the query for searching the agent profile
type GetAgentProfileQuery struct {
	APIKey string
}

// ReturnedAgentProfile defines the returned agent profile
type ReturnedAgentProfile struct {
	ID                string
	InitialEmotion    *emotion.Emotion   `gorm:"embedded"`
	Personality       *ocean.Personality `gorm:"embedded"`
	IsContextInfinite bool
	ContextExistsTime *time.Duration `gorm:"embedded"`
}
