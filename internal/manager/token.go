package manager

import (
	"fmt"
	"time"
)

// GetEmotionToken gets the token for getting emotion
func GetEmotionToken(
	agentKey string,
	contextID string,
) string {
	now := time.Now().Nanosecond()
	return fmt.Sprintf(
		"%s:%s_%s_%d",
		"emotion",
		agentKey,
		contextID,
		now,
	)
}
