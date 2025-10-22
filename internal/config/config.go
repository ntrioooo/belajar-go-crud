// internal/config/config.go
package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port   string
	Secret string
	DSN    string
}

func Load() *Config {
	port := getEnv("PORT", "8080")
	secret := getEnv("SECRET", "changeme")

	// Prioritas 1: DATABASE_URL (format DSN lengkap)
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return &Config{Port: port, Secret: secret, DSN: dsn}
	}

	// Prioritas 2: rakit dari DB_* (cocok untuk docker compose)
	host := getEnv("DB_HOST", "belajar-go-db") // <— penting
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASS", "postgres")
	name := getEnv("DB_NAME", "postgres")
	dbPort := getEnv("DB_PORT", "5432")
	ssl := getEnv("DB_SSLMODE", "disable")
	tz := getEnv("DB_TIMEZONE", "Asia/Jakarta")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, pass, name, dbPort, ssl, tz,
	)
	return &Config{Port: port, Secret: secret, DSN: dsn}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
