package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"monopc-starter/internal/app/config" // Sesuaikan path ini
	"strconv"
	"time"
)

type DB struct {
	*gorm.DB
}

func NewPostgresGormDB(cfg *config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	sqlDB, sdErr := gormDB.DB()
	if sdErr != nil {
		return nil, sdErr
	}

	maxOpenConns, mocErr := strconv.Atoi(cfg.MaxOpenConns)
	if mocErr != nil {
		return nil, mocErr
	}

	maxConnLifetime, mclErr := time.ParseDuration(cfg.MaxConnLifetime)
	if mclErr != nil {
		return nil, mclErr
	}

	maxIdleLifetime, milErr := time.ParseDuration(cfg.MaxIdleLifetime)
	if milErr != nil {
		return nil, milErr
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(maxConnLifetime)
	sqlDB.SetConnMaxIdleTime(maxIdleLifetime)

	return gormDB, err
}
