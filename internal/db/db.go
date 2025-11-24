package db

import (
	"fmt"

	"github.com/iamasocial/hightalent-test-task/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB creates a new GORM Postgres connection
func NewPostgresDB(cfg config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := db.AutoMigrate(&Question{}, &Answer{}); err != nil {
		return nil, fmt.Errorf("automigration failed: %w", err)
	}

	return db, nil
}
