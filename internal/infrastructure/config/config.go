package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Server      struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	JWT struct {
		AccessSecret   string
		RefreshSecret  string
		AccessExpires  string
		RefreshExpires string
	}
}

func LoadConfig() *Config {
	_ = godotenv.Load() // carrega .env se existir

	cfg := &Config{}

	cfg.Environment = getEnv("APP_ENV", "development")

	// Server
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "123456")
	cfg.Database.DBName = getEnv("DB_NAME", "rental_db")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT
	cfg.JWT.AccessSecret = getEnv(
		"JWT_ACCESS_SECRET",
		"super-access-secret-key",
	)

	cfg.JWT.RefreshSecret = getEnv(
		"JWT_REFRESH_SECRET",
		"super-refresh-secret-key",
	)

	cfg.JWT.AccessExpires = getEnv(
		"JWT_ACCESS_EXPIRES",
		"15m",
	)

	cfg.JWT.RefreshExpires = getEnv(
		"JWT_REFRESH_EXPIRES",
		"168h",
	)

	return cfg
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
