package main

import (
	"context"
	"log"
	"yosem/internal/config"
	"yosem/internal/lib/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
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

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
