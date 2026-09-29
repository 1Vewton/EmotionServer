package database

import (
	"context"
	"fmt"

	"github.com/1Vewton/EmotionServer/internal/ocean"
	"github.com/1Vewton/EmotionServer/internal/profile"
	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"gorm.io/gorm"
)

// DB defines the connection to the database
var DB *gorm.DB

// Connect connects to the database
func Connect(
	databaseURL string,
	databaseType databasetype.DatabaseType,
	tables ...any,
) (*gorm.DB, error) {
	driver, err := databaseType.GetDriver(databaseURL)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(
		driver,
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	// Create Tables
	err = db.AutoMigrate(tables...)
	return db, err
}

// Close closes the connection
func Close(
	db *gorm.DB,
) error {
	sql, err := db.DB()
	if err != nil {
		return err
	}
	err = sql.Close()
	return err
}

// AddNewAgentProfile creates new profile
func AddNewAgentProfile(
	ctx context.Context,
	db *gorm.DB,
	personality *ocean.Personality,
	apiKey string,
) error {
	newAgentProfile, err := profile.NewAgentProfile(
		apiKey,
		personality,
	)
	if err != nil {
		return err
	}
	// Checks if this agent already exists
	result, err := gorm.G[profile.AgentProfile](db).Where(
		&profile.AgentProfile{
			APIKey: apiKey,
		},
	).Find(ctx)
	if err != nil {
		return err
	}
	if len(result) > 0 {
		return fmt.Errorf(
			"agent with api %s already exists",
			apiKey,
		)
	}
	// Create
	err = gorm.G[profile.AgentProfile](db).Create(
		ctx,
		newAgentProfile,
	)
	return err
}

// SearchAgentProfile searches profile for agent
func SearchAgentProfile(
	db *gorm.DB,
	apiKey string,
) *profile.AgentProfile {
	fetchedProfile := &profile.AgentProfile{}
	db.Where(
		profile.GetSearchProfileByAPIKeySentence(apiKey),
	).First(&fetchedProfile)
	return fetchedProfile
}
