package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
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
		Secret string
	}
}

func LoadConfig() *Config {
	_ = godotenv.Load() // carrega .env se existir

	cfg := &Config{}

	// Server
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "senha123")
	cfg.Database.DBName = getEnv("DB_NAME", "rental_db")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT
	cfg.JWT.Secret = getEnv("JWT_SECRET", "sua-chave-secreta-muito-forte")

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
