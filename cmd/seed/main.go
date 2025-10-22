// main.go (ringkas, fokus ke DB init + seed)
package main

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"belajar-go/internal/core/domain"
	gormrepo "belajar-go/internal/repository/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL tidak di-set. Set env ini di docker-compose.yml (gunakan host=belajar-go-db).")
	}

	db := mustOpenWithRetry(dsn)

	if err := db.AutoMigrate(&gormrepo.User{}); err != nil {
		log.Fatal(err)
	}

	seedAdmin(db)
	log.Println("✅ Ready.")
}

func mustOpenWithRetry(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error

	delay := 500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err == nil {
			sqlDB, derr := db.DB()
			if derr == nil && sqlDB.PingContext(ctx) == nil {
				sqlDB.SetMaxOpenConns(10)
				sqlDB.SetMaxIdleConns(5)
				sqlDB.SetConnMaxLifetime(30 * time.Minute)
				return db
			}
		}
		if ctx.Err() != nil {
			log.Fatalf("DB connect timeout: %v", err)
		}
		time.Sleep(delay)
		if delay < 4*time.Second {
			delay *= 2
		}
	}
}

func seedAdmin(db *gorm.DB) {
	email := "admin@admin.com"
	var count int64
	if err := db.Model(&gormrepo.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		log.Println("admin already exists, skip seeding.")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("4dm1n@123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	admin := gormrepo.User{
		Email:    email,
		Username: "admin",
		Password: string(hash),
		Role:     domain.RoleAdmin,
	}
	if err := db.Create(&admin).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("✅ Admin user seeded:", email)
}
