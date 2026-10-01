package manager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/redis/go-redis/v9"
)

// ContextManager defines the manager for context
type ContextManager struct {
	client         *redis.Client
	profileManager *AgentProfileManager
}

// NewContextManager creates new context manager
func NewContextManager(
	client *redis.Client,
	profileManager *AgentProfileManager,
) *ContextManager {
	return &ContextManager{
		client:         client,
		profileManager: profileManager,
	}
}

// NewContext creates new context
// Return the context key
func (manager *ContextManager) NewContext(
	ctx context.Context,
	apiKey string,
	contextID string,
) (string, error) {
	agentProfile, err := manager.profileManager.SearchAgentProfile(
		ctx,
		apiKey,
	)
	if err != nil {
		return "", err
	}
	storedProfile := agentProfile.ToStoredProfile()
	stringResult, err := json.Marshal(
		storedProfile,
	)
	if err != nil {
		return "", err
	}
	// Adds new context
	key := GetEmotionToken(
		apiKey,
		contextID,
	)
	exists, err := manager.client.Exists(
		ctx,
		key,
	).Result()
	if err != nil {
		return "", err
	}
	if exists == 0 {
		_, err = manager.client.Set(
			ctx,
			key,
			stringResult,
			0,
		).Result()
		return key, err
	}
	return key, fmt.Errorf(
		"context key of %s already exists",
		key,
	)
}

// GetDataFromContext gets the stored profile from context key
func (manager *ContextManager) GetDataFromContext(
	ctx context.Context,
	contextKey string,
) (*profile.StoredProfile, error) {
	var result *profile.StoredProfile
	stringResult, err := manager.client.Get(
		ctx,
		contextKey,
	).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(
		[]byte(stringResult),
		&result,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
