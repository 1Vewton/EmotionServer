package emotion

import (
	"errors"
	"fmt"

	"github.com/1Vewton/EmotionServer/pkg/mathematics"
)

// Emotion defines the emotion
type Emotion struct {
	Pleasure  float64 `json:"pleasure"`
	Arousal   float64 `json:"arousal"`
	Dominance float64 `json:"dominance"`
	Certainty float64 `json:"certainty"`
	Novelty   float64 `json:"novelty"`
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
		Novelty:   novelty,
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
		emotion.Novelty,
	)
}

// Equals see if two emotions are equal
func (emotion *Emotion) Equals(
	emotion2 *Emotion,
) bool {
	if mathematics.RoundDigits(
		emotion.Pleasure,
		0.01,
	) == mathematics.RoundDigits(
		emotion2.Pleasure,
		0.01,
	) && mathematics.RoundDigits(
		emotion.Arousal,
		0.01,
	) == mathematics.RoundDigits(
		emotion2.Arousal,
		0.01,
	) && mathematics.RoundDigits(
		emotion.Dominance,
		0.01,
	) == mathematics.RoundDigits(
		emotion2.Dominance,
		0.01,
	) && mathematics.RoundDigits(
		emotion.Certainty,
		0.01,
	) == mathematics.RoundDigits(
		emotion2.Certainty,
		0.01,
	) && mathematics.RoundDigits(
		emotion.Novelty,
		0.01,
	) == mathematics.RoundDigits(
		emotion2.Novelty,
		0.01,
	) {
		return true
	}
	return false
}

// Copy copies emotion
func (emotion *Emotion) Copy() *Emotion {
	return &Emotion{
		Pleasure:  emotion.Pleasure,
		Arousal:   emotion.Arousal,
		Dominance: emotion.Dominance,
		Certainty: emotion.Certainty,
		Novelty:   emotion.Novelty,
	}
}
