package database

import (
	"github.com/redis/go-redis/v9"
)

// RedisClient defines the client for redis
var RedisClient *redis.Client

// CloseRedis closes the redis client
func CloseRedis(
	client *redis.Client,
) error {
	err := client.Close()
	return err
}
