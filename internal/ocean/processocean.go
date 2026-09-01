package ocean

import (
	"github.com/1Vewton/EmotionServer/internal/emotion"
)

// GetPleasure gets the pleasure of the initial emotion.
func (personality *Personality) GetPleasure() float64 {
	processMatirx := GetOceanLineOne()
	return personality.DotProductProcessing(
		[5]float64(processMatirx),
	)
}

// GetArousal gets the arousal of the initial emotion.
func (personality *Personality) GetArousal() float64 {
	processMatirx := GetOceanLineTwo()
	return personality.DotProductProcessing(
		[5]float64(processMatirx),
	)
}

// GetDominance gets the dominance of the initial emotion.
func (personality *Personality) GetDominance() float64 {
	processMatirx := GetOceanLineThree()
	return personality.DotProductProcessing(
		[5]float64(processMatirx),
	)
}

// GetCertainty gets the certainty of the initial emotion
func (personality *Personality) GetCertainty() float64 {
	processMatirx := GetOceanLineFour()
	return personality.DotProductProcessing(
		[5]float64(processMatirx),
	)
}

// GetNovelty gets the novelty of the initial emotion
func (personality *Personality) GetNovelty() float64 {
	processMatirx := GetOceanLineFive()
	return personality.DotProductProcessing(
		[5]float64(processMatirx),
	)
}

// GetInitialEmotion gets the initial emotion according to the ocean personality
func (personality *Personality) GetInitialEmotion() (*emotion.Emotion, error) {
	return emotion.NewEmotion(
		personality.GetPleasure(),
		personality.GetArousal(),
		personality.GetDominance(),
		personality.GetCertainty(),
		personality.GetNovelty(),
	)
}
