package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"gorm.io/gorm"
)

// AgentProfileManager defines the manager for agent profile
type AgentProfileManager struct {
	db *gorm.DB
}

// NewAgentProfileManager creates new agent profile manager
func NewAgentProfileManager(
	db *gorm.DB,
) *AgentProfileManager {
	return &AgentProfileManager{
		db: db,
	}
}

// AddNewAgentProfile creates new profile
func (manager *AgentProfileManager) AddNewAgentProfile(
	ctx context.Context,
	newAgentProfile *profile.AgentProfile,
) error {
	if newAgentProfile == nil {
		return errors.New(
			"you cannot pass a nil profile",
		)
	}
	// Checks if this agent already exists
	result, err := gorm.G[profile.AgentProfile](manager.db).Where(
		&profile.AgentProfile{
			APIKey: newAgentProfile.APIKey,
		},
	).Find(ctx)
	if err != nil {
		return err
	}
	if len(result) > 0 {
		return fmt.Errorf(
			"agent with api %s already exists",
			newAgentProfile.APIKey,
		)
	}
	// Create
	err = gorm.G[profile.AgentProfile](manager.db).Create(
		ctx,
		newAgentProfile,
	)
	return err
}

// SearchAgentProfile searches profile for agent
func (manager *AgentProfileManager) SearchAgentProfile(
	ctx context.Context,
	apiKey string,
) (*profile.AgentProfile, error) {
	fetchedProfile, err := gorm.G[*profile.AgentProfile](manager.db).Where(
		&profile.AgentProfile{
			APIKey: apiKey,
		},
	).First(ctx)
	return fetchedProfile, err
}
