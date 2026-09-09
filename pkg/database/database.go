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
) error {
	driver, err := databaseType.GetDriver(databaseURL)
	if err != nil {
		return err
	}
	DB, err = gorm.Open(
		driver,
		&gorm.Config{},
	)
	if err != nil {
		return err
	}
	// Create Tables
	err = DB.AutoMigrate(tables...)
	return err
}

// Close closes the connection
func Close() {
	sql, err := DB.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = sql.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// AddNewAgentProfile creates new profile
func AddNewAgentProfile(
	ctx context.Context,
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
	result, err := gorm.G[profile.AgentProfile](DB).Where(
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
	err = gorm.G[profile.AgentProfile](DB).Create(
		ctx,
		newAgentProfile,
	)
	return err
}
