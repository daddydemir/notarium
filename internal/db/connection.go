package db

import (
	"fmt"

	"github.com/daddydemir/notarium/internal/domain"
	"github.com/daddydemir/notarium/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *utils.Config) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMmode,
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed: %w", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.Entry{},
		&domain.Tag{},
		&domain.Topic{},
		&domain.TopicTag{},
		&domain.Note{},
		&domain.File{},
		&domain.Reminder{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
