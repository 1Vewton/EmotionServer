package database

import (
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
