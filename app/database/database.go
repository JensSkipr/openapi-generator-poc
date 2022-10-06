/* This file is auto-generated, manual edits in this file will be overwritten! */
package database

import (
	"errors"

	"github.com/nightborn-be/blink/skipr-test/app/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func checkConfig(config map[string]string) error {
	hasConnectionString := false

	// Check if config has keys
	for key := range config {
		if key == "ConnectionString" {
			hasConnectionString = true
		}
	}

	if !hasConnectionString {
		return errors.New(MISSING_CONFIG_PARAMS + "ConnectionString")
	}

	return nil
}

func Initialise(config map[string]string) (*gorm.DB, error) {
	// Check config
	if err := checkConfig(config); err != nil {
		return nil, err
	}

	var (
		db  *gorm.DB
		err error
	)

	if dsn := config["ConnectionString"]; dsn != "" {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("local.db"))
	}
	if err != nil {
		return nil, errors.New(FAILED_CONNECTION)
	}

	// Migrate the schema
	if err := db.AutoMigrate(models.Expense{}); err != nil {
		return nil, err
	}

	return db, nil
}

func InitialiseTest() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, errors.New(FAILED_CONNECTION)
	}

	// Migrate the schema
	if err := db.AutoMigrate(models.Expense{}); err != nil {
		return nil, err
	}

	// Seed
	if err := Seed(db); err != nil {
		return nil, err
	}

	return db, nil
}

func InitialiseTestFromTestId(testId string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(testId + ".db"))
	if err != nil {
		return nil, errors.New(FAILED_CONNECTION)
	}

	// Migrate the schema
	if err := db.AutoMigrate(models.Expense{}); err != nil {
		return nil, err
	}

	// Seed
	if err := Seed(db); err != nil {
		return nil, err
	}

	return db, nil
}