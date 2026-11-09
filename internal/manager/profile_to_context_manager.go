package manager

import (
	"github.com/redis/go-redis/v9"
)

// ProfileToContextManager defines the manager that manages contexts for profile
type ProfileToContextManager struct {
	client         *redis.Client
	profileManager *AgentProfileManager
}

// NewProfileToContextManager creates new ProfileToContextManager
func NewProfileToContextManager(
	client *redis.Client,
	profileManager *AgentProfileManager,
) *ProfileToContextManager {
	return &ProfileToContextManager{
		client:         client,
		profileManager: profileManager,
	}
}

// AddContext adds new context for agent profile
func (manager *ProfileToContextManager) AddContext(
	apiKey string,
	contextKey string,
) error {
	return nil
}
