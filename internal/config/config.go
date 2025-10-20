package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Secret string
	DSN    string
}

func Load() *Config {
	_ = godotenv.Load(".env") // optional: tidak fatal kalau tidak ada

	c := &Config{
		Port:   getenv("PORT", "8080"),
		Secret: mustGetenv("SECRET"),
		DSN:    getenv("DATABASE_DSN", "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"),
	}
	return c
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing required env: %s", k)
	}
	return v
}
