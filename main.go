package main

import (
	"context"
	"log"
	"yosem/internal/config"
	"yosem/internal/lib/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	log := logger.SetupLogger(cfg.Env)
	_ = log

	dsn := "postgres://postgres:123@localhost:5433"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Error("Not connected to database")
	}
	defer pool.Close()
}
