package manager

import (
	"testing"
	"time"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/google/uuid"
)

// TestEmotionCRUD tests the crud of emotion
func TestEmotionCRUD(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	testClient := database.NewRedisConfig(
		"localhost:6379",
		"",
	).WithDialTimeout(
		5,
	).WithReadTimeout(
		5,
	).WithMaxRetries(
		3,
	).WithMaxRetryBackoff(
		5,
	).WithMinRetryBackoff(
		3,
	).ToClient()
	// Define data
	agentKey := uuid.NewString()
	contextKey := uuid.NewString()
	testProfile, err := profile.NewAgentProfile(
		agentKey,
		ocean.NewPersonality(
			0.0,
			0.0,
			0.0,
			0.0,
			0.0,
		),
		5*time.Hour,
	)
	if err != nil {
		t.Fatal(err)
	}
	// Define manager
	emotionManager := NewEmotionManager(
		testClient,
	)
	err = emotionManager.AddEmotionInfo(
		ctx,
		agentKey,
		contextKey,
		testProfile,
		5*time.Hour,
	)
	if err != nil {
		t.Fatal(err)
	}
	isContextExists, err := emotionManager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isContextExists {
		t.Fatalf(
			"emotion context with key %s does not exists",
			contextKey,
		)
	}
}
