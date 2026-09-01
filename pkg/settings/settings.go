package settings

import (
	"fmt"

	"github.com/1Vewton/EmotionServer/pkg/databasetype"
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	serverPort   *string
	serverHost   *string
	databaseURL  *string
	databaseType *int
}

// Initialize reads the env file setted
func (s *settings) Initialize(filePath string) error {
	err := godotenv.Load(filePath)
	return err
}

// GetServerPort gets the port of the service running on
func (s *settings) GetServerPort() string {
	return SetConfigString(
		"SERVER_PORT",
		"3392",
		&s.serverPort,
	)
}

// GetServerHost gets the host of the service running on
func (s *settings) GetServerHost() string {
	return SetConfigString(
		"SERVER_HOST",
		"0.0.0.0",
		&s.serverHost,
	)
}

// GetDatabaseURL gets the url of the database
func (s *settings) GetDatabaseURL() string {
	return SetConfigString(
		"DATABASE_URL",
		"file::memory:?cache=shared",
		&s.databaseURL,
	)
}

// GetDatabaseType gets the type of the database
func (s *settings) GetDatabaseType() databasetype.DatabaseType {
	return databasetype.ToDatabaseType(
		SetConfigInteger(
			"DATABASE_TYPE",
			0,
			&s.databaseType,
		),
	)
}

// GetServerURL gets the url of the service
func (s *settings) GetServerURL() string {
	return fmt.Sprintf(
		"%s:%s",
		s.GetServerHost(),
		s.GetServerPort(),
	)
}
