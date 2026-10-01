package database

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig defines the config for redis
type RedisConfig struct {
	RedisOptions *redis.Options
}

// NewRedisConfig creates new redis config
func NewRedisConfig(
	address string,
	password string,
) *RedisConfig {
	return &RedisConfig{
		RedisOptions: &redis.Options{
			Addr:     address,
			Password: password,
			DB:       0,
			Protocol: 2,
		},
	}
}

// ToClient converts config to client
func (cfg *RedisConfig) ToClient() *redis.Client {
	return redis.NewClient(
		cfg.RedisOptions,
	)
}

// WithDialTimeout sets dial timeout
func (cfg *RedisConfig) WithDialTimeout(
	dialTimeout int,
) *RedisConfig {
	cfg.RedisOptions.DialTimeout = time.Duration(dialTimeout) * time.Second
	return cfg
}

// WithReadTimeout sets read timeout
func (cfg *RedisConfig) WithReadTimeout(
	readTimeout int,
) *RedisConfig {
	cfg.RedisOptions.ReadTimeout = time.Duration(readTimeout) * time.Second
	return cfg
}

// WithWriteTimeout sets write timeout
func (cfg *RedisConfig) WithWriteTimeout(
	writeTimeout int,
) *RedisConfig {
	cfg.RedisOptions.WriteTimeout = time.Duration(writeTimeout) * time.Second
	return cfg
}

// WithMaxRetries sets max retries
func (cfg *RedisConfig) WithMaxRetries(
	maxRetries int,
) *RedisConfig {
	cfg.RedisOptions.MaxRetries = maxRetries
	return cfg
}

// WithMaxRetryBackoff sets max retry backoff
func (cfg *RedisConfig) WithMaxRetryBackoff(
	maxRetryBackoff int,
) *RedisConfig {
	cfg.RedisOptions.MaxRetryBackoff = time.Duration(maxRetryBackoff) * time.Second
	return cfg
}

// WithMinRetryBackoff sets max retry backoff
func (cfg *RedisConfig) WithMinRetryBackoff(
	minRetryBackoff int,
) *RedisConfig {
	cfg.RedisOptions.MinRetryBackoff = time.Duration(minRetryBackoff) * time.Second
	return cfg
}
