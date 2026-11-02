package manager

import (
	"context"
	"encoding/json"
	"errors"
	"time"

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

// ContextExistsWithAgentProfile checks if certain agent profile has context
// system: agentprofile -> context
func (manager *ContextManager) ContextExistsWithAgentProfile(
	ctx context.Context,
	apiKey string,
	contextKey string,
) (bool, error) {
	keyExists, err := manager.client.Exists(
		ctx,
		apiKey,
	).Result()
	if err != nil {
		return false, err
	}
	if keyExists < 0 {
		return false, nil
	}
	existsInProfileList, err := manager.client.SIsMember(
		ctx,
		apiKey,
		contextKey,
	).Result()
	if err != nil {
		return false, err
	}
	return existsInProfileList, err
}

// AddNewContextForProfile adds new context under an agent profile
// system: agentprofile -> context
func (manager *ContextManager) AddNewContextForProfile(
	ctx context.Context,
	apiKey string,
	contextKey string,
) error {
	exists, err := manager.ContextExistsWithAgentProfile(
		ctx,
		apiKey,
		contextKey,
	)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	_, err = manager.client.SAdd(
		ctx,
		apiKey,
		contextKey,
	).Result()
	return err
}

// NewContext creates new context
// Return the context key
func (manager *ContextManager) NewContext(
	ctx context.Context,
	apiKey string,
	contextID string,
	orderedContextExistsTime *time.Duration,
) (string, error) {
	agentProfile, err := manager.profileManager.SearchAgentProfile(
		ctx,
		apiKey,
	)
	if err != nil {
		return "", err
	}
	var contextExistsTime time.Duration
	if orderedContextExistsTime != nil {
		contextExistsTime = *orderedContextExistsTime
	} else if agentProfile.IsContextInfinite {
		contextExistsTime = 0
	} else {
		if agentProfile.ContextExistsTime != nil {
			contextExistsTime = *agentProfile.ContextExistsTime
		} else {
			return "", errors.New(
				"no context exists time is given",
			)
		}
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
			contextExistsTime,
		).Result()
		manager.AddNewContextForProfile(
			ctx,
			apiKey,
			key,
		)
		return key, err
	}
	return key, nil
}

// HasContext checks if certain context exists
func (manager *ContextManager) HasContext(
	ctx context.Context,
	contextKey string,
) (bool, error) {
	exists, err := manager.client.Exists(
		ctx,
		contextKey,
	).Result()
	if err != nil {
		return false, err
	}
	if exists == 0 {
		return false, nil
	}
	return true, nil
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
