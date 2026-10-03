package manager

import (
	"testing"
	"time"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
)

// TestContextCRUD tests the crud of the context
func TestContextCRUD(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	// connect database
	tDB, err := database.Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Error(err)
	}
	newAgentProfileManager := NewAgentProfileManager(
		tDB,
	)
	mutualPersonality := ocean.NewPersonality(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	apiKey := "test114514"
	newProfile, err := profile.NewAgentProfile(
		apiKey,
		mutualPersonality,
		5*time.Minute,
	)
	err = newAgentProfileManager.AddNewAgentProfile(
		ctx,
		newProfile,
	)
	if err != nil {
		t.Error(err)
	}
	// connect to the redis
	resultClient := database.NewRedisConfig(
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
	newContextManager := NewContextManager(
		resultClient,
		newAgentProfileManager,
	)
	// Search
	ctxID := "ctx114514"
	key, err := newContextManager.NewContext(
		ctx,
		apiKey,
		ctxID,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	exists, err := newContextManager.HasContext(
		ctx,
		key,
	)
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Errorf(
			"it is suggested that %s does not exists or an error occurs in HasContext",
			key,
		)
	}
	resultData, err := newContextManager.GetDataFromContext(
		ctx,
		key,
	)
	if err != nil {
		t.Error(err)
	}
	expectedInitialEmotion, err := mutualPersonality.GetInitialEmotion()
	if err != nil {
		t.Error(err)
	}
	if !resultData.InitialEmotion.Equals(
		expectedInitialEmotion,
	) {
		t.Errorf(
			"the result data is inconsistent wih expected data",
		)
	}
	err = database.CloseRedis(resultClient)
	if err != nil {
		t.Error(err)
	}
	err = database.Close(tDB)
	if err != nil {
		t.Error(err)
	}
}
