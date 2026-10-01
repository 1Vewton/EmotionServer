package database

import (
	"testing"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
)

// TestRedisConfig tests config for redis
func TestRedisConfig(t *testing.T) {
	t.Parallel()
	resultConfig := NewRedisConfig(
		"localhost:6379",
		"",
	).WithMaxRetries(
		3,
	)
	resultMaxRetries := resultConfig.RedisOptions.MaxRetries
	if resultMaxRetries != 3 {
		t.Errorf(
			"expected %d, got %d",
			3,
			resultMaxRetries,
		)
	}
}

// TestContextCRUD tests the crud of the context
func TestContextCRUD(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	// connect database
	tDB, err := Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	mutualPersonality := ocean.NewPersonality(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	apiKey := "test114514"
	err = AddNewAgentProfile(
		ctx,
		tDB,
		mutualPersonality,
		apiKey,
	)
	if err != nil {
		t.Error(err)
	}
	// connect to the redis
	resultClient := NewRedisConfig(
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
	// Search
	ctxID := "ctx114514"
	key, err := NewContext(
		ctx,
		resultClient,
		tDB,
		apiKey,
		ctxID,
	)
	if err != nil {
		t.Error(err)
	}
	resultData, err := GetDataFromContext(
		ctx,
		resultClient,
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
	err = CloseRedis(resultClient)
	if err != nil {
		t.Error(err)
	}
	err = Close(tDB)
	if err != nil {
		t.Error(err)
	}
}
