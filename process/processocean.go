package process

import (
	"github.com/1Vewton/EmotionServer/emotion"
	"github.com/1Vewton/EmotionServer/ocean"
	"github.com/1Vewton/EmotionServer/utils/mathematics"
)

// GetPleasure gets the pleasure of the initial emotion.
func GetPleasure(personality ocean.Personality) float64 {
	processMatirx := ocean.GetOceanLineOne()
	return mathematics.DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetArousal gets the arousal of the initial emotion.
func GetArousal(personality ocean.Personality) float64 {
	processMatirx := ocean.GetOceanLineTwo()
	return mathematics.DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetDominance gets the dominance of the initial emotion.
func GetDominance(personality ocean.Personality) float64 {
	processMatirx := ocean.GetOceanLineThree()
	return mathematics.DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetCertainty gets the certainty of the initial emotion
func GetCertainty(personality ocean.Personality) float64 {
	processMatirx := ocean.GetOceanLineFour()
	return mathematics.DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetNovelty gets the novelty of the initial emotion
func GetNovelty(personality ocean.Personality) float64 {
	processMatirx := ocean.GetOceanLineFive()
	return mathematics.DotProductProcessing(
		personality,
		[5]float64(processMatirx),
	)
}

// GetInitialEmotion gets the initial emotion according to the ocean personality
func GetInitialEmotion(
	personality ocean.Personality,
) (*emotion.Emotion, error) {
	return emotion.NewEmotion(
		GetPleasure(personality),
		GetArousal(personality),
		GetDominance(personality),
		GetCertainty(personality),
		GetNovelty(personality),
	)
}
