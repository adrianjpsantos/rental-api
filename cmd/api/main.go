package main

import (
	"fmt"
	"log"

	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/database"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/routes"
)

func main() {
	// Carrega configurações
	cfg := config.LoadConfig()

	// Conecta com o banco
	fmt.Println("🚀 Conectando Aos Bancos de Dados...")
	db, err := database.NewConnectionPostgres(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	fmt.Println("🚀 Iniciando Rental API...")
	port := cfg.Server.Port
	// Aqui vamos configurar o router depois
	router := routes.SetupRouter(db)

	log.Printf("✅ API rodando na porta %s", port)

	if err := router.Listen(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
