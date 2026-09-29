package profile

import (
	"time"

	"github.com/1Vewton/EmotionServer/internal/emotion"
)

// StoredProfile defines the profile stored in redis
type StoredProfile struct {
	LastUpdatedTime time.Time        `json:"last_updated_time"`
	InitialEmotion  *emotion.Emotion `json:"initial_emotion"`
	CurrentEmotion  *emotion.Emotion `json:"curent_emotion"`
}

// NewStoredProfile creates new stored profile
func NewStoredProfile(
	InitialEmotion *emotion.Emotion,
) *StoredProfile {
	return &StoredProfile{
		LastUpdatedTime: time.Now(),
		InitialEmotion:  InitialEmotion,
		CurrentEmotion: InitialEmotion,
	}
}
