package mathematics

import (
	"github.com/1Vewton/EmotionServer/ocean"
)

// DotProductProcessing gets the result of processing personality
func DotProductProcessing(
	personality ocean.Personality,
	processMatrix [5]float64,
) float64 {
	result := 0.0
	result += personality.Openness * processMatrix[0]
	result += personality.Conscientiousness * processMatrix[1]
	result += personality.Agreeableness * processMatrix[2]
	result += personality.Extraversion * processMatrix[3]
	result += personality.Neuroticism * processMatrix[4]
	return max(min(result, 0.40), -0.40)
}
