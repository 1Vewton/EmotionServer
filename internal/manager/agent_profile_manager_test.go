package manager

import (
	"testing"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/database"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/google/uuid"
)

// TestProfileFetching tests the fetching for profile
func TestProfileFetching(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	id := uuid.NewString()
	tDB, err := database.Connect(
		":memory:",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Error(err)
	}
	newAgentProfileManager := NewAgentProfileManager(
		tDB,
	)
	mutualPersonality := ocean.NewPersonality(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	err = newAgentProfileManager.AddNewAgentProfile(
		ctx,
		mutualPersonality,
		id,
	)
	if err != nil {
		t.Error(err)
	}
	result, err := newAgentProfileManager.SearchAgentProfile(
		ctx,
		id,
	)
	if err != nil {
		t.Error(err)
	}
	if !result.Personality.Equals(result.Personality) {
		t.Error(
			"the result fetched from two datas are not the same",
		)
	}
	err = database.Close(tDB)
	if err != nil {
		t.Error(err)
	}
}
