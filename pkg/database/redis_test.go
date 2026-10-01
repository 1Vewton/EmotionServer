package database

import (
	"testing"
)

// TestRedisConfig tests config for redis
func TestRedisConfig(t *testing.T) {
	t.Parallel()
	resultConfig := NewRedisConfig(
		"localhost:6379",
		"",
	).WithMaxRetries(
		3,
	)
	resultMaxRetries := resultConfig.RedisOptions.MaxRetries
	if resultMaxRetries != 3 {
		t.Errorf(
			"expected %d, got %d",
			3,
			resultMaxRetries,
		)
	}
}
