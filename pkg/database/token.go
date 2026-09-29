package database

import (
	"fmt"
)

// GetEmotionToken gets the token for getting emotion
func GetEmotionToken(
	apiKey string,
	contextID string,
) string {
	return fmt.Sprintf(
		"%s:%s_%s",
		"emotion",
		apiKey,
		contextID,
	)
}
