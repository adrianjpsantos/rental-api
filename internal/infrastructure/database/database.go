package database

import (
	"fmt"
	"log"

	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionPostgreSql(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, // melhora performance
	})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco: %w", err)
	}

	log.Println("✅ Conectado ao PostgreSQL com sucesso!")
	return db, nil
}
