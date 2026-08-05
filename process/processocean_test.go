package process

import (
	"testing"

	"github.com/1Vewton/EmotionServer/emotion"
	"github.com/1Vewton/EmotionServer/ocean"
)

// TestGetInitialEmotionMutual tests the initial emotion
func TestGetInitialEmotionMutual(
	t *testing.T,
) {
	mutualPersonality := ocean.NewPersonality(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	targetEmotion, errEmotion := emotion.NewEmotion(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	if errEmotion != nil {
		t.Error(errEmotion.Error())
	}
	initialEmotionMutual, errInitialEmotion := GetInitialEmotion(
		*mutualPersonality,
	)
	if errInitialEmotion != nil {
		t.Error(errInitialEmotion.Error())
	}
	// Check data
	if *initialEmotionMutual != *targetEmotion {
		t.Errorf(
			"Expected %s, got %s",
			targetEmotion.ShowEmotionInfo(),
			initialEmotionMutual.ShowEmotionInfo(),
		)
	}
}
