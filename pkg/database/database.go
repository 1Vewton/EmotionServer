package database

import (
	"fmt"

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
