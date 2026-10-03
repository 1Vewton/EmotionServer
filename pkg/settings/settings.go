package settings

import (
	"fmt"
	"time"

	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	serverPort            *string
	serverHost            *string
	databaseURL           *string
	databaseType          *int
	redisURL              *string
	redisPassword         *string
	redisDialTimeout      *int
	redisReadTimeout      *int
	redisWriteTimeout     *int
	redisMaxRetries       *int
	redisMinRetryBackoff  *int
	redisMaxRetryBackoff  *int
	contextLastTimeInDays *int
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
func (s *settings) GetRedisURL() string {
	return SetConfigString(
		"REDIS_URL",
		"localhost:6379",
		&s.redisURL,
	)
}

// GetRedisPassword gets the url for the Redis
func (s *settings) GetRedisPassword() string {
	return SetConfigString(
		"REDIS_PASSWORD",
		"",
		&s.redisPassword,
	)
}

// GetRedisDialTimeout gets the Dial Timeout for the Redis
func (s *settings) GetRedisDialTimeout() int {
	return SetConfigInteger(
		"REDIS_DIAL_TIMEOUT",
		10,
		&s.redisDialTimeout,
	)
}

// GetRedisReadTimeout gets the Read Timeout for the Redis
func (s *settings) GetRedisReadTimeout() int {
	return SetConfigInteger(
		"REDIS_READ_TIMEOUT",
		5,
		&s.redisReadTimeout,
	)
}

// GetRedisWriteTimeout gets the Write Timeout for the Redis
func (s *settings) GetRedisWriteTimeout() int {
	return SetConfigInteger(
		"REDIS_WRITE_TIMEOUT",
		5,
		&s.redisWriteTimeout,
	)
}

// GetRedisMaxRetries gets the max retries for the Redis
func (s *settings) GetRedisMaxRetries() int {
	return SetConfigInteger(
		"REDIS_MAX_RETRIES",
		5,
		&s.redisMaxRetries,
	)
}

// GetRedisMaxRetryBackOff gets the max retry backoff for the Redis
func (s *settings) GetRedisMaxRetryBackOff() int {
	return SetConfigInteger(
		"REDIS_MAX_RETRY_BACKOFF",
		100,
		&s.redisMaxRetryBackoff,
	)
}

// GetRedisMinRetryBackOff gets the min retry backoff for the Redis
func (s *settings) GetRedisMinRetryBackOff() int {
	return SetConfigInteger(
		"REDIS_MIN_RETRY_BACKOFF",
		10,
		&s.redisMinRetryBackoff,
	)
}

// GetContextLastTimeInDays gets the last time for context in days
func (s *settings) GetContextLastTimeInDays() time.Duration {
	daysNum := SetConfigInteger(
		"CONTEXT_LAST_TIME_IN_DAYS",
		7,
		&s.redisMinRetryBackoff,
	)
	return time.Duration(daysNum) * 24 * time.Hour
}
