package settings

import (
	"fmt"

	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	serverPort           *string
	serverHost           *string
	databaseURL          *string
	databaseType         *int
	redisURL             *string
	redisPassword        *string
	redisDialTimeout     *int
	redisReadTimeout     *int
	redisWriteTimeout    *int
	redisMaxRetries      *int
	redisMinRetryBackoff *int
	redisMaxRetryBackoff *int
}

// Initialize reads the env file setted
func (s *settings) Initialize(filePath string) error {
	err := godotenv.Load(filePath)
	return err
}

// GetServerPort gets the port of the service running on
func (s *settings) GetServerPort() string {
	return SetConfigString(
		"SERVER_PORT",
		"3392",
		&s.serverPort,
	)
}

// GetServerHost gets the host of the service running on
func (s *settings) GetServerHost() string {
	return SetConfigString(
		"SERVER_HOST",
		"0.0.0.0",
		&s.serverHost,
	)
}

// GetDatabaseURL gets the url of the database
func (s *settings) GetDatabaseURL() string {
	return SetConfigString(
		"DATABASE_URL",
		"file::memory:?cache=shared",
		&s.databaseURL,
	)
}

// GetDatabaseType gets the type of the database
func (s *settings) GetDatabaseType() databasetype.DatabaseType {
	return databasetype.ToDatabaseType(
		SetConfigInteger(
			"DATABASE_TYPE",
			0,
			&s.databaseType,
		),
	)
}

// GetServerURL gets the url of the service
func (s *settings) GetServerURL() string {
	return fmt.Sprintf(
		"%s:%s",
		s.GetServerHost(),
		s.GetServerPort(),
	)
}

// GetRedisURL gets the url for the Redis
func (cfg *settings) GetRedisURL() string {
	return SetConfigString(
		"REDIS_URL",
		"localhost:6379",
		&cfg.redisURL,
	)
}

// GetRedisPassword gets the url for the Redis
func (cfg *settings) GetRedisPassword() string {
	return SetConfigString(
		"REDIS_PASSWORD",
		"",
		&cfg.redisPassword,
	)
}

// GetRedisDialTimeout gets the Dial Timeout for the Redis
func (cfg *settings) GetRedisDialTimeout() int {
	return SetConfigInteger(
		"REDIS_DIAL_TIMEOUT",
		10,
		&cfg.redisDialTimeout,
	)
}

// GetRedisReadTimeout gets the Read Timeout for the Redis
func (cfg *settings) GetRedisReadTimeout() int {
	return SetConfigInteger(
		"REDIS_READ_TIMEOUT",
		5,
		&cfg.redisReadTimeout,
	)
}

// GetRedisWriteTimeout gets the Write Timeout for the Redis
func (cfg *settings) GetRedisWriteTimeout() int {
	return SetConfigInteger(
		"REDIS_WRITE_TIMEOUT",
		5,
		&cfg.redisWriteTimeout,
	)
}

// GetRedisMaxRetries gets the max retries for the Redis
func (cfg *settings) GetRedisMaxRetries() int {
	return SetConfigInteger(
		"REDIS_MAX_RETRIES",
		5,
		&cfg.redisMaxRetries,
	)
}

// GetRedisMaxRetryBackOff gets the max retry backoff for the Redis
func (cfg *settings) GetRedisMaxRetryBackOff() int {
	return SetConfigInteger(
		"REDIS_MAX_RETRY_BACKOFF",
		100,
		&cfg.redisMaxRetryBackoff,
	)
}

// GetRedisMinRetryBackOff gets the min retry backoff for the Redis
func (cfg *settings) GetRedisMinRetryBackOff() int {
	return SetConfigInteger(
		"REDIS_MIN_RETRY_BACKOFF",
		10,
		&cfg.redisMinRetryBackoff,
	)
}
