package profile

import (
	"github.com/1Vewton/EmotionServer/internal/emotion"
	"github.com/1Vewton/EmotionServer/internal/ocean"
	"gorm.io/gorm"
)

// AgentProfile defines the profile for user to store in the database
type AgentProfile struct {
	gorm.Model
	ID             string
	APIKey         string            `gorm:"unique"`
	InitialEmotion emotion.Emotion   `gorm:"embedded"`
	Personality    ocean.Personality `gorm:"embedded"`
}
