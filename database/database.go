package database

import (
	"fmt"
	"github.com/dicapisar/13cc_manager/config"
	"github.com/dicapisar/13cc_manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
}

type DatabaseImpl struct {
	DB *gorm.DB
}

func NewDatabase(config *config.DatabaseConfig) (Database, error) {
	dns := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.User, config.Password, config.Name, config.Host, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if config.AutoMigrate {
		if err := autoMigrate(db); err != nil {
			return nil, err
		}
	}

	return &DatabaseImpl{
		DB: db,
	}, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Rol{},
		&models.UserRol{},
	)
}
