package database

import (
	"context"
	"encoding/json"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// RedisClient defines the client for redis
var RedisClient *redis.Client

// NewContext creates new context
func NewContext(
	ctx context.Context,
	client *redis.Client,
	db *gorm.DB,
	apiKey string,
	contextID string,
) (string, error) {
	agentProfile, err := SearchAgentProfile(
		ctx,
		db,
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
	exists, err := client.Exists(
		ctx,
		key,
	).Result()
	if err != nil {
		return "", err
	}
	if exists == 0 {
		_, err = client.Set(
			ctx,
			key,
			stringResult,
			0,
		).Result()
		return "", err
	}
	return key, nil
}

// GetDataFromContext
func GetDataFromContext(
	ctx context.Context,
	client *redis.Client,
	contextID string,
) (*profile.StoredProfile, error) {
	var result *profile.StoredProfile
	stringResult, err := client.Get(
		ctx,
		contextID,
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

// CloseRedis closes the redis client
func CloseRedis(
	client *redis.Client,
) error {
	err := client.Close()
	return err
}
