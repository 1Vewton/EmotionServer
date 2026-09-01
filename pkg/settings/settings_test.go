package settings

import (
	"testing"

	"github.com/1Vewton/EmotionServer/pkg/database"
)

// Test the env reading of string var
func TestEnvReadingString(t *testing.T) {
	res := GetEnvString("test", "T")
	if res != "T" {
		t.Errorf("Expected T, got %s", res)
	}
	t.Setenv("test", "KK3S")
	res = GetEnvString("test", "T")
	if res != "KK3S" {
		t.Errorf("Expected KK3S, got %s", res)
	}
}

// Test the env setting
func TestEnvSettingString(t *testing.T) {
	var tSetting *settings = &settings{}
	t.Setenv("SERVER_PORT", "114514")
	res := tSetting.GetServerPort()
	if res != "114514" {
		t.Errorf("Expected 114514, got %s", res)
	}
}

// TestEnvReadingInteger tests the env reading of integer var
func TestEnvReadingInteger(t *testing.T) {
	res := GetEnvInteger("TEST", 0)
	if res != 0 {
		t.Errorf(
			"Expected %d, got %d",
			0,
			res,
		)
	}
	t.Setenv("test", "1")
	res = GetEnvInteger("TEST", 0)
	if res != 1 {
		t.Errorf(
			"Expected %d, got %d",
			1,
			res,
		)
	}
}

// TestEnvSettingDatabaseType tests the env setting of database type
func TestEnvSettingDatabaseType(t *testing.T) {
	t.Setenv("DATABASE_Type", "0")
	var tSetting *settings = &settings{}
	res := tSetting.GetDatabaseType()
	if res != database.Sqlite {
		t.Errorf(
			"Expected %d, got %d",
			database.Sqlite,
			res,
		)
	}
}
