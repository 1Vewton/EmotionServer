package manager

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ProfileToContextManager defines the manager that manages contexts for profile
type ProfileToContextManager struct {
	client *redis.Client
}

// NewProfileToContextManager creates new ProfileToContextManager
func NewProfileToContextManager(
	client *redis.Client,
) *ProfileToContextManager {
	return &ProfileToContextManager{
		client: client,
	}
}

// AddContext adds new context for agent profile
func (manager *ProfileToContextManager) AddContext(
	ctx context.Context,
	agentKey string,
	contextKey string,
) error {
	contextExists, err := manager.ContextExistsWithAgentProfile(
		ctx,
		agentKey,
		contextKey,
	)
	if err != nil {
		return err
	}
	if contextExists {
		return nil
	}
	_, err = manager.client.SAdd(
		ctx,
		agentKey,
		contextKey,
	).Result()
	return err
}

// ContextExistsWithProfile checks whether a context exists in profile and the profile exists
func (manager *ProfileToContextManager) ContextExistsWithAgentProfile(
	ctx context.Context,
	agentKey string,
	contextKey string,
) (bool, error) {
	profileExists, err := manager.ProfileExists(
		ctx,
		agentKey,
	)
	if err != nil {
		return false, err
	}
	if !profileExists {
		return false, nil
	}
	return manager.client.SIsMember(
		ctx,
		agentKey,
		contextKey,
	).Result()
}

// ProfileExists checks if a profile exists
func (manager *ProfileToContextManager) ProfileExists(
	ctx context.Context,
	agentKey string,
) (bool, error) {
	intResult, err := manager.client.Exists(
		ctx,
		agentKey,
	).Result()
	if err != nil {
		return false, err
	}
	if intResult == 1 {
		return true, nil
	}
	return false, nil
}

// GetContextsInProfile gets contexts stored in agent profile
func (manager *ProfileToContextManager) GetContextsInProfile(
	ctx context.Context,
	agentKey string,
) ([]string, error) {
	hasProfile, err := manager.ProfileExists(
		ctx,
		agentKey,
	)
	if err != nil {
		return nil, err
	}
	if !hasProfile {
		return nil, nil
	}
	contexts, err := manager.client.SMembers(
		ctx,
		agentKey,
	).Result()
	return contexts, err
}

// DeleteContextFromProfile deletes context from profile
func (manager *ProfileToContextManager) DeleteContextFromProfile(
	ctx context.Context,
	agentKey string,
	contextKey string,
) error {
	contextExistsWithProfile, err := manager.ContextExistsWithAgentProfile(
		ctx,
		agentKey,
		contextKey,
	)
	if err != nil {
		return err
	}
	if !contextExistsWithProfile {
		return nil
	}
	_, err = manager.client.SRem(
		ctx,
		agentKey,
		contextKey,
	).Result()
	return err
}
