package ocean

import (
	"testing"

	"github.com/1Vewton/EmotionServer/internal/emotion"
)

// TestGetInitialEmotionMutual tests the initial emotion for mutual personality
func TestGetInitialEmotionMutual(
	t *testing.T,
) {
	mutualPersonality := NewPersonality(
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
	if !initialEmotionMutual.Equals(targetEmotion) {
		t.Errorf(
			"Expected %s, got %s",
			targetEmotion.ShowEmotionInfo(),
			initialEmotionMutual.ShowEmotionInfo(),
		)
	}
}

// TestGetInitialEmotionHelper tests the initial emotion for helper emotion
func TestGetInitialEmotionHelper(
	t *testing.T,
) {
	helperPersonality := NewPersonality(
		0.30,
		0.50,
		0.40,
		0.80,
		-0.20,
	)
	targetEmotion, errEmotion := emotion.NewEmotion(
		0.38,
		0.24,
		0.06,
		0.33,
		0.02,
	)
	if errEmotion != nil {
		t.Error(errEmotion.Error())
	}
	initialEmotionHelper, errInitialEmotion := GetInitialEmotion(
		*helperPersonality,
	)
	if errInitialEmotion != nil {
		t.Error(errInitialEmotion.Error())
	}
	// Check data
	if !initialEmotionHelper.Equals(targetEmotion) {
		t.Errorf(
			"Expected %s, got %s",
			targetEmotion.ShowEmotionInfo(),
			initialEmotionHelper.ShowEmotionInfo(),
		)
	}
}

// TestGetInitialEmotionAnalyst tests the initial emotion for analyst emotion
func TestGetInitialEmotionAnalyst(
	t *testing.T,
) {
	analystPersonality := NewPersonality(
		0.40,
		0.80,
		-0.30,
		0.10,
		-0.40,
	)
	targetEmotion, errEmotion := emotion.NewEmotion(
		-0.02,
		0.11,
		0.05,
		0.30,
		0.00,
	)
	if errEmotion != nil {
		t.Error(errEmotion.Error())
	}
	initialEmotionAnalyst, errInitialEmotion := GetInitialEmotion(
		*analystPersonality,
	)
	if errInitialEmotion != nil {
		t.Error(errInitialEmotion.Error())
	}
	// Check data
	if !initialEmotionAnalyst.Equals(targetEmotion) {
		t.Errorf(
			"Expected %s, got %s",
			targetEmotion.ShowEmotionInfo(),
			initialEmotionAnalyst.ShowEmotionInfo(),
		)
	}
}
