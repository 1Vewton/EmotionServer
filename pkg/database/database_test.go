package database

import (
	"testing"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/google/uuid"
)

// TestConnection tests the connection
func TestConnection(t *testing.T) {
	t.Parallel()
	tDB, err := Connect(
		"file::memory:?cache=shared",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Error(err)
	}
	err = Close(tDB)
	if err != nil {
		t.Error(err)
	}
}

// TestProfileFetching tests the fetching for profile
func TestProfileFetching(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	id := uuid.NewString()
	tDB, err := Connect(
		"file::memory:?cache=shared",
		databasetype.Sqlite,
		&profile.AgentProfile{},
	)
	if err != nil {
		t.Error(err)
	}
	mutualPersonality := ocean.NewPersonality(
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)
	err = AddNewAgentProfile(
		ctx,
		tDB,
		mutualPersonality,
		id,
	)
	if err != nil {
		t.Error(err)
	}
	result := SearchAgentProfile(
		tDB,
		id,
	)
	if !result.Personality.Equals(result.Personality) {
		t.Error(
			"the result fetched from two datas are not the same",
		)
	}
	err = Close(tDB)
	if err != nil {
		t.Error(err)
	}
}
