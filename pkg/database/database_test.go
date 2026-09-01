package database

import (
	"testing"

	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
)

// TestConnection tests the connection
func TestConnection(t *testing.T) {
	t.Parallel()
	err := Connect(
		"file::memory:?cache=shared",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Error(err)
	}
	Close()
}
