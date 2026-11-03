package manager

import (
	"testing"
	"time"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/google/uuid"
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
	if err != nil {
		t.Fatal(err)
	}
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
	err = newContextManager.DeleteContext(
		ctx,
		key,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = newContextManager.DeleteContextFromProfile(
		ctx,
		apiKey,
		key,
	)
	if err != nil {
		t.Fatal(err)
	}
	existsAfterDelete, err := newContextManager.HasContext(
		ctx,
		key,
	)
	if err != nil {
		t.Fatal(err)
	}
	if existsAfterDelete {
		t.Fatalf(
			"expected %s not to exists",
			key,
		)
	}
	err = newContextManager.DeleteContextFromProfile(
		ctx,
		apiKey,
		key,
	)
	if err != nil {
		t.Fatal(err)
	}
	contextExistsWithProfile, err := newContextManager.ContextExistsWithAgentProfile(
		ctx,
		apiKey,
		key,
	)
	if err != nil {
		t.Fatal(err)
	}
	if contextExistsWithProfile {
		t.Fatalf(
			"expected context key of %s does not exists",
			key,
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

// TestContextWithAgentProfile checks combination of context and agent profile
func TestContextWithAgentProfile(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	// connect to the redis
	apiKey := uuid.NewString()
	contextID := uuid.NewString()
	redisclient := database.NewRedisConfig(
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
	tDB, err := database.Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Fatal(err)
	}
	testAgentProfile := NewAgentProfileManager(
		tDB,
	)
	newAgentProfile, err := profile.NewAgentProfile(
		apiKey,
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
	newContextManager := NewContextManager(
		redisclient,
		testAgentProfile,
	)
	err = testAgentProfile.AddNewAgentProfile(
		ctx,
		newAgentProfile,
	)
	if err != nil {
		t.Fatal(err)
	}
	contextKey, err := newContextManager.NewContext(
		ctx,
		apiKey,
		contextID,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = newContextManager.AddNewContextForProfile(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	exists, err := newContextManager.ContextExistsWithAgentProfile(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf(
			"context with key %s expected to exists",
			contextKey,
		)
	}
	contextExists, err := newContextManager.HasContext(
		ctx,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !contextExists {
		t.Fatalf(
			"this context %s just not exists",
			contextKey,
		)
	}
	err = database.CloseRedis(redisclient)
	if err != nil {
		t.Error(err)
	}
	err = database.Close(tDB)
	if err != nil {
		t.Error(err)
	}
}

// TestDataCleaning tests the data cleaning
func TestDataCleaning(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	apiKey := uuid.NewString()
	contextID := uuid.NewString()
	tDB, err := database.Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Fatal(err)
	}
	redisclient := database.NewRedisConfig(
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
	profileManager := NewAgentProfileManager(
		tDB,
	)
	testContextManager := NewContextManager(
		redisclient,
		profileManager,
	)
	newProfile, err := profile.NewAgentProfile(
		apiKey,
		ocean.NewPersonality(
			0.0,
			0.0,
			0.0,
			0.0,
			0.0,
		),
		5*time.Hour,
	)
	err = profileManager.AddNewAgentProfile(
		ctx,
		newProfile,
	)
	if err != nil {
		t.Fatal(err)
	}
	contextKey, err := testContextManager.NewContext(
		ctx,
		apiKey,
		contextID,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testContextManager.AddNewContextForProfile(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testContextManager.CleanSingleData(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	hasContext1, err := testContextManager.HasContext(
		ctx,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !hasContext1 {
		t.Fatalf(
			"expected context key %s to exist",
			contextKey,
		)
	}
	hasContextInProfile1, err := testContextManager.ContextExistsWithAgentProfile(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !hasContextInProfile1 {
		t.Fatalf(
			"expected context with %s exists for profile %s",
			contextKey,
			apiKey,
		)
	}
	newContextKey := uuid.NewString()
	err = testContextManager.AddNewContextForProfile(
		ctx,
		apiKey,
		newContextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testContextManager.CleanSingleData(
		ctx,
		apiKey,
		newContextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	hasContextInProfile2, err := testContextManager.ContextExistsWithAgentProfile(
		ctx,
		apiKey,
		newContextKey,
	)
	if err != nil {
		t.Fatal(err)
	}
	if hasContextInProfile2 {
		t.Fatalf(
			"expected context with key %s to get deleted",
			newContextKey,
		)
	}
}
