package database

import (
	"database/sql"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	_ "github.com/lib/pq"
)

func NewConnectionPostgres(cfg *config.Config) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// testa conexão de verdade
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL successfully.")

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return nil, err
	}

	fmt.Println("Postgres version:", version)

	return db, nil
}
