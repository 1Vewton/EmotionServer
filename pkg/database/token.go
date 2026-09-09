package database

import (
	"fmt"
)

// GetEmotionToken gets the token for getting emotion
func GetEmotionToken(
	apiKey string,
) string {
	return fmt.Sprintf(
		"%s:%s",
		"emotion",
		apiKey,
	)
}
