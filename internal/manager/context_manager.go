package manager

import (
	"context"
	"errors"
	"time"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/redis/go-redis/v9"
)

// ContextManager defines the manager for context
type ContextManager struct {
	client                  *redis.Client
	profileManager          *AgentProfileManager
	profileToContextManager *ProfileToContextManager
	emotionManager          *EmotionManager
}

// NewContextManager creates new context manager
func NewContextManager(
	client *redis.Client,
	profileManager *AgentProfileManager,
) *ContextManager {
	profileToContextManager := NewProfileToContextManager(
		client,
	)
	emotionManager := NewEmotionManager(
		client,
	)
	return &ContextManager{
		client:                  client,
		profileManager:          profileManager,
		profileToContextManager: profileToContextManager,
		emotionManager:          emotionManager,
	}
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
	// AddNewContext
	contextKey := GetEmotionToken(
		agentProfile.AgentKey,
		contextID,
	)
	err = manager.emotionManager.AddEmotionInfo(
		ctx,
		agentProfile.AgentKey,
		contextKey,
		agentProfile,
		contextExistsTime,
	)
	if err != nil {
		return contextKey, err
	}
	err = manager.profileToContextManager.AddContext(
		ctx,
		agentKey,
		contextKey,
	)
	return contextKey, err
}

// HasContext checks if certain context exists
func (manager *ContextManager) HasContext(
	ctx context.Context,
	agentKey string,
	contextKey string,
) (bool, error) {
	contextExists, err := manager.emotionManager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		return false, err
	}
	contextExistsInProfile, err := manager.profileToContextManager.ContextExistsWithAgentProfile(
		ctx,
		agentKey,
		contextKey,
	)
	return contextExists && contextExistsInProfile, err
}

// GetEmotionFromContext gets the stored profile from context key
func (manager *ContextManager) GetEmotionFromContext(
	ctx context.Context,
	contextKey string,
) (*profile.StoredProfile, error) {
	return manager.emotionManager.GetEmotionFromContext(
		ctx,
		contextKey,
	)
}

// DeleteContext deletes context
func (manager *ContextManager) DeleteContext(
	ctx context.Context,
	contextKey string,
) error {
	exists, err := manager.HasContext(
		ctx,
		"",
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
	return manager.profileToContextManager.ProfileExists(
		ctx,
		agentKey,
	)
}

// GetContextsInProfile gets all the context keys in a collection of agent profile
func (manager *ContextManager) GetContextsInProfile(
	ctx context.Context,
	agentKey string,
) ([]string, error) {
	return manager.profileToContextManager.GetContextsInProfile(
		ctx,
		agentKey,
	)
}

// DeleteProfileData deletes a profile stored in redis and the data of its context
func (manager *ContextManager) DeleteProfileData(
	ctx context.Context,
	agentKey string,
) error {
	contexts, err := manager.profileToContextManager.GetContextsInProfile(
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
	contextExists, err := manager.emotionManager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		return err
	}
	contextExistsWithProfile, err := manager.profileToContextManager.ContextExistsWithAgentProfile(
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
	contexts, err := manager.profileToContextManager.GetContextsInProfile(
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
