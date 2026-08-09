package settings

import (
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	serverPort *string
	serverURL  *string
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

// GetServerURL gets the url of the service running on
func (s *settings) GetServerURL() string {
	return SetConfigString(
		"SERVER_URL",
		"0.0.0.0",
		&s.serverURL,
	)
}
