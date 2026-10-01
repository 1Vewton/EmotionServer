package database

import (
	"fmt"
	"time"
)

// GetEmotionToken gets the token for getting emotion
func GetEmotionToken(
	apiKey string,
	contextID string,
) string {
	now := time.Now().Nanosecond()
	return fmt.Sprintf(
		"%s:%s_%s_%d",
		"emotion",
		apiKey,
		contextID,
		now,
	)
}
