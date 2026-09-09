package database

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient defines the client for redis
var RedisClient *redis.Client

// InitRedisClient inits the redis client
func InitRedisClient(
	address string,
	password string,
	dialTimeout int,
	readTimeout int,
	writeTimeout int,
	maxRetries int,
	maxRetryBackoff int,
	minRetryBackoff int,
) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:            address,
			Password:        password,
			DB:              0,
			Protocol:        2,
			DialTimeout:     time.Duration(dialTimeout) * time.Second,
			ReadTimeout:     time.Duration(readTimeout) * time.Second,
			WriteTimeout:    time.Duration(writeTimeout) * time.Second,
			MaxRetries:      maxRetries,
			MaxRetryBackoff: time.Duration(maxRetryBackoff) * time.Millisecond,
			MinRetryBackoff: time.Duration(minRetryBackoff) * time.Millisecond,
		},
	)
}

// CloseRedis closes the redis client
func CloseRedis(
	client *redis.Client,
) error {
	err := client.Close()
	return err
}
