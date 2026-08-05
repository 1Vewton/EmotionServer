package emotion

import (
	"errors"
	"fmt"
)

// Emotion defines the emotion
type Emotion struct {
	Pleasure  float64
	Arousal   float64
	Dominance float64
	Certainty float64
	Novalty   float64
}

// NewEmotion creates new PADCN emotion status
func NewEmotion(
	pleasure float64,
	arousal float64,
	dominance float64,
	certainty float64,
	novelty float64,
) (*Emotion, error) {
	if pleasure < -1.0 || pleasure > 1.0 ||
		arousal < -1.0 || arousal > 1.0 ||
		dominance < -1.0 || dominance > 1.0 ||
		certainty < -1.0 || certainty > 1.0 ||
		novelty < -1.0 || novelty > 1.0 {
		return nil, errors.New(
			"All the values should be between -1.0 and 1.0",
		)
	}
	return &Emotion{
		Pleasure:  pleasure,
		Arousal:   arousal,
		Dominance: dominance,
		Certainty: certainty,
		Novalty:   novelty,
	}, nil
}

// ShowEmotionInfo shows the info of the emotion
func (emotion *Emotion) ShowEmotionInfo() string {
	return fmt.Sprintf(
		"P:%f;A:%f,D:%f;C:%f;N:%f",
		emotion.Pleasure,
		emotion.Arousal,
		emotion.Dominance,
		emotion.Certainty,
		emotion.Novalty,
	)
}
