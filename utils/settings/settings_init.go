//go:build !test

package settings

import (
	"fmt"

	"github.com/1Vewton/EmotionServer/utils/logger"
)

// Read the .env file
func init() {
	err := Settings.Initialize(".env")
	if err != nil {
		logger.SysLogger.Error(
			fmt.Sprintf(
				"Env reading failed due to %s",
				err,
			),
		)
	}
}
