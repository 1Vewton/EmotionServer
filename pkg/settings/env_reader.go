package settings

import (
	"os"
	"strconv"

	"github.com/1Vewton/EmotionServer/pkg/logger"
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

// GetEnvInteger returns the int value stored in the environment
func GetEnvInteger(
	key string,
	defaultValue int,
) int {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	num, errTurnInteger := strconv.Atoi(result)
	if errTurnInteger != nil {
		logger.SysLogger.Error(errTurnInteger.Error())
		panic(errTurnInteger.Error())
	}
	return num
}
