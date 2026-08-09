package settings

import (
	"os"
)

// GetEnvString returns the string value stored in the environment
func GetEnvString(
	key string,
	defaultValue string,
) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
