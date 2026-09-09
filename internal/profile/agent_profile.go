package profile

import (
	"github.com/1Vewton/EmotionServer/internal/emotion"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AgentProfile defines the profile for user to store in the database
type AgentProfile struct {
	gorm.Model
	ID             string
	APIKey         string             `gorm:"unique"`
	InitialEmotion *emotion.Emotion   `gorm:"embedded"`
	Personality    *ocean.Personality `gorm:"embedded"`
}

// NewAgentProfile creates new agent profile
func NewAgentProfile(
	apiKey string,
	personality *ocean.Personality,
) (*AgentProfile, error) {
	id := uuid.NewString()
	initialEmotion, err := personality.GetInitialEmotion()
	if err != nil {
		return nil, err
	}
	return &AgentProfile{
		ID:             id,
		APIKey:         apiKey,
		InitialEmotion: initialEmotion,
		Personality:    personality,
	}, nil
}
