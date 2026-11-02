package emotionapi

import (
	"github.com/1Vewton/EmotionServer/internal/emotion"
)

// NewContextQuery defines parameter for query that creates new context
type NewContextQuery struct {
	APIKey string
}

// WithContextResponse returns response with context key
type WithContextResponse struct {
	ContextKey string
}

// GetCurrentEmotionQuery defines parameter for getting current emotion
type GetCurrentEmotionQuery struct {
	APIKey     string
	ContextKey string
}

// GetCurrentEmotionResponse defines the response for getting current emotion
type GetCurrentEmotionResponse struct {
	CurrentEmotion *emotion.Emotion
	NewContextKey  *string
}
