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
	port := getEnv("PORT", "3001")
	secret := getEnv("SECRET", "changeme")

	// 1) DATABASE_URL kalau ada, dipakai duluan (baik lokal maupun docker)
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return &Config{Port: port, Secret: secret, DSN: dsn}
	}

	// 2) Rakitan dari DB_*
	host := getEnv("DB_HOST", defaultDBHost())
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASS", "password")
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

// defaultDBHost: kalau jalan DI DALAM Docker, pakai nama service; kalau TIDAK, pakai localhost.
func defaultDBHost() string {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "belajar-go-db" // dalam container
	}
	// Boleh juga override manual pakai ENV RUN_MODE=local/docker
	if os.Getenv("RUN_MODE") == "docker" {
		return "belajar-go-db"
	}
	return "localhost" // proses lokal
}
