package ocean

import (
	"github.com/1Vewton/EmotionServer/internal/emotion"
)

// GetPleasure gets the pleasure of the initial emotion.
func GetPleasure(personality Personality) float64 {
	processMatirx := GetOceanLineOne()
	return DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetArousal gets the arousal of the initial emotion.
func GetArousal(personality Personality) float64 {
	processMatirx := GetOceanLineTwo()
	return DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetDominance gets the dominance of the initial emotion.
func GetDominance(personality Personality) float64 {
	processMatirx := GetOceanLineThree()
	return DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetCertainty gets the certainty of the initial emotion
func GetCertainty(personality Personality) float64 {
	processMatirx := GetOceanLineFour()
	return DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetNovelty gets the novelty of the initial emotion
func GetNovelty(personality Personality) float64 {
	processMatirx := GetOceanLineFive()
	return DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetInitialEmotion gets the initial emotion according to the ocean personality
func GetInitialEmotion(
	personality Personality,
) (*emotion.Emotion, error) {
	return emotion.NewEmotion(
		GetPleasure(personality),
		GetArousal(personality),
		GetDominance(personality),
		GetCertainty(personality),
		GetNovelty(personality),
	)
}
