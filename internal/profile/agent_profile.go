package profile

import (
	"time"

	"github.com/1Vewton/EmotionServer/internal/emotion"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AgentProfile defines the profile for user to store in the database
type AgentProfile struct {
	gorm.Model
	ID                string
	APIKey            string             `gorm:"unique"`
	InitialEmotion    *emotion.Emotion   `gorm:"embedded"`
	Personality       *ocean.Personality `gorm:"embedded"`
	IsContextInfinite bool
	ContextExistsTime *time.Duration
}

// Equals tests if two profiles are equal
func (profile *AgentProfile) Equals(
	another *AgentProfile,
) bool {
	return profile.InitialEmotion.Equals(another.InitialEmotion) &&
		profile.Personality.Equals(another.Personality) &&
		profile.ID == another.ID &&
		profile.APIKey == another.APIKey
}

// ToStoredProfile converts AgentProfile to StoredProfile
func (profile *AgentProfile) ToStoredProfile() *StoredProfile {
	return &StoredProfile{
		LastUpdatedTime: time.Now(),
		InitialEmotion:  profile.InitialEmotion.Copy(),
		CurrentEmotion:  profile.InitialEmotion.Copy(),
	}
}

// SetIsContextInfinite sets whether the context is inifinite
func (profile *AgentProfile) SetIsContextInfinite(
	isInfinite bool,
) *AgentProfile {
	var infDuration time.Duration = 0
	if isInfinite {
		profile.ContextExistsTime = &infDuration
	}
	profile.IsContextInfinite = isInfinite
	return profile
}

// NewAgentProfile creates new agent profile
func NewAgentProfile(
	apiKey string,
	personality *ocean.Personality,
	defaultTime time.Duration,
) (*AgentProfile, error) {
	id := uuid.NewString()
	initialEmotion, err := personality.GetInitialEmotion()
	if err != nil {
		return nil, err
	}
	return &AgentProfile{
		ID:                id,
		APIKey:            apiKey,
		InitialEmotion:    initialEmotion,
		Personality:       personality,
		IsContextInfinite: false,
		ContextExistsTime: &defaultTime,
	}, nil
}

// GetSearchProfileByAPIKeySentence gets the sentence for searching agent profile
func GetSearchProfileByAPIKeySentence(
	apiKey string,
) *AgentProfile {
	return &AgentProfile{
		APIKey: apiKey,
	}
}
