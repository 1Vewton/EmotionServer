package manager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/redis/go-redis/v9"
)

// ContextManager defines the manager for context
type ContextManager struct {
	client                  *redis.Client
	profileManager          *AgentProfileManager
	profileToContextManager *ProfileToContextManager
}

// NewContextManager creates new context manager
func NewContextManager(
	client *redis.Client,
	profileManager *AgentProfileManager,
) *ContextManager {
	profileToContextManager := NewProfileToContextManager(
		client,
		profileManager,
	)
	return &ContextManager{
		client:                  client,
		profileManager:          profileManager,
		profileToContextManager: profileToContextManager,
	}
}

// ContextExistsWithAgentProfile checks if certain agent profile has context
// system: agentprofile -> context
func (manager *ContextManager) ContextExistsWithAgentProfile(
	ctx context.Context,
	agentKey string,
	contextKey string,
) (bool, error) {
	keyExists, err := manager.client.Exists(
		ctx,
		agentKey,
	).Result()
	if err != nil {
		return false, err
	}
	if keyExists < 0 {
		return false, nil
	}
	existsInProfileList, err := manager.client.SIsMember(
		ctx,
		agentKey,
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
	agentKey string,
	contextKey string,
) error {
	exists, err := manager.ContextExistsWithAgentProfile(
		ctx,
		agentKey,
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
		agentKey,
		contextKey,
	).Result()
	return err
}

// NewContext creates new context
// Return the context key
func (manager *ContextManager) NewContext(
	ctx context.Context,
	agentKey string,
	contextID string,
	orderedContextExistsTime *time.Duration,
) (string, error) {
	agentProfile, err := manager.profileManager.SearchAgentProfile(
		ctx,
		agentKey,
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
		agentKey,
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
			agentKey,
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

// DeleteContext deletes context
func (manager *ContextManager) DeleteContext(
	ctx context.Context,
	contextKey string,
) error {
	exists, err := manager.HasContext(
		ctx,
		contextKey,
	)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	_, err = manager.client.Del(
		ctx,
		contextKey,
	).Result()
	return err
}

// DeleteContextFromProfile deletes the context from the context list of the profile
func (manager *ContextManager) DeleteContextFromProfile(
	ctx context.Context,
	agentKey string,
	contextKey string,
) error {
	hasProfile, err := manager.HasProfile(
		ctx,
		agentKey,
	)
	if err != nil {
		return err
	}
	if !hasProfile {
		return nil
	}
	_, err = manager.client.SRem(
		ctx,
		agentKey,
		contextKey,
	).Result()
	return err
}

// HasProfile checks if certain profile exists in the redis database
func (manager *ContextManager) HasProfile(
	ctx context.Context,
	agentKey string,
) (bool, error) {
	exists, err := manager.client.Exists(
		ctx,
		agentKey,
	).Result()
	if err != nil {
		return false, err
	}
	if exists == 0 {
		return false, nil
	}
	return true, nil
}

// GetContextsInProfile gets all the context keys in a collection of agent profile
func (manager *ContextManager) GetContextsInProfile(
	ctx context.Context,
	agentKey string,
) ([]string, error) {
	hasProfile, err := manager.HasProfile(
		ctx,
		agentKey,
	)
	if err != nil {
		return nil, err
	}
	if !hasProfile {
		return nil, fmt.Errorf(
			"profile with api-key %s does not exists",
			agentKey,
		)
	}
	contexts, err := manager.client.SMembers(
		ctx,
		agentKey,
	).Result()
	return contexts, err
}

// DeleteProfileData deletes a profile stored in redis and the data of its context
func (manager *ContextManager) DeleteProfileData(
	ctx context.Context,
	agentKey string,
) error {
	contexts, err := manager.GetContextsInProfile(
		ctx,
		agentKey,
	)
	if err != nil {
		return err
	}
	for _, selectedContext := range contexts {
		manager.DeleteContext(
			ctx,
			selectedContext,
		)
	}
	_, err = manager.client.Del(
		ctx,
		agentKey,
	).Result()
	if err != nil {
		return err
	}
	return nil
}

// CleanSingleData cleans single data
func (manager *ContextManager) CleanSingleData(
	ctx context.Context,
	agentKey string,
	contextKey string,
) error {
	contextExists, err := manager.HasContext(
		ctx,
		contextKey,
	)
	if err != nil {
		return err
	}
	contextExistsWithProfile, err := manager.ContextExistsWithAgentProfile(
		ctx,
		agentKey,
		contextKey,
	)
	if err != nil {
		return err
	}
	if contextExistsWithProfile && !contextExists {
		err = manager.DeleteContextFromProfile(
			ctx,
			agentKey,
			contextKey,
		)
		return err
	}
	return nil
}

// CleanAllData cleans all contexts for a profile
func (manager *ContextManager) CleanAllData(
	ctx context.Context,
	agentKey string,
) error {
	contexts, err := manager.GetContextsInProfile(
		ctx,
		agentKey,
	)
	if err != nil {
		return err
	}
	for _, singleContextKey := range contexts {
		err := manager.CleanSingleData(
			ctx,
			agentKey,
			singleContextKey,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
