package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/redis/go-redis/v9"
)

// EmotionManager manages emotions for contexts
type EmotionManager struct {
	client *redis.Client
}

// NewEmotionManager creates new emotion manager
func NewEmotionManager(
	client *redis.Client,
) *EmotionManager {
	return &EmotionManager{
		client: client,
	}
}

// ContextExists checks if certain context exists
func (manager *EmotionManager) ContextExists(
	ctx context.Context,
	contextKey string,
) (bool, error) {
	existsData, err := manager.client.Exists(
		ctx,
		contextKey,
	).Result()
	if err != nil {
		return false, err
	}
	if existsData == 1 {
		return true, nil
	}
	return false, nil
}

// AddEmotionInfo adds the emotion of an agent to the database.
// The emotion will consist of the initial emotion, current emotion and time of the last update
func (manager *EmotionManager) AddEmotionInfo(
	ctx context.Context,
	agentKey string,
	contextKey string,
	agentProfile *profile.AgentProfile,
	contextExistsTime time.Duration,
) error {
	exists, err := manager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	storedProfile := agentProfile.ToStoredProfile()
	bytesData, err := json.Marshal(storedProfile)
	if err != nil {
		return err
	}
	_, err = manager.client.Set(
		ctx,
		contextKey,
		string(bytesData),
		contextExistsTime,
	).Result()
	return err
}

// GetEmotionFromContext gets emotion from context
func (manager *EmotionManager) GetEmotionFromContext(
	ctx context.Context,
	contextKey string,
) (*profile.StoredProfile, error) {
	contextExists, err := manager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		return nil, err
	}
	if !contextExists {
		return nil, fmt.Errorf(
			"context with key %s does not exists",
			contextKey,
		)
	}
	stringEmotion, err := manager.client.Get(
		ctx,
		contextKey,
	).Result()
	if err != nil {
		return nil, err
	}
	bytesEmotion := []byte(stringEmotion)
	var result *profile.StoredProfile
	err = json.Unmarshal(
		bytesEmotion,
		&result,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteEmotion deletes emotion
func (manager *EmotionManager) DeleteEmotion(
	ctx context.Context,
	contextKey string,
) error {
	contextExists, err := manager.ContextExists(
		ctx,
		contextKey,
	)
	if err != nil {
		return err
	}
	if contextExists {
		return nil
	}
	_, err = manager.client.Del(
		ctx,
		contextKey,
	).Result()
	return err
}
