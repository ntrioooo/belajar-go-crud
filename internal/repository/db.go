// internal/repository/db.go
package repository

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := pingWithRetry(sqlDB, 30*time.Second); err != nil {
		return nil, err
	}
	return db, nil
}

func pingWithRetry(db *sql.DB, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	delay := 500 * time.Millisecond
	for {
		if err := db.Ping(); err == nil {
			return nil
		}
		if time.Now().After(deadline) {
			return db.Ping()
		}
		time.Sleep(delay)
		if delay < 4*time.Second {
			delay *= 2
		}
	}
}
