package main

import (
	"fmt"
	"os"

	"logsence/internal/config"
	"logsence/internal/handler"
	"logsence/pkg/loger"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		panic(fmt.Errorf("failed to load conifg: %w", err))
	}

	lg := loger.New(cfg.Log.Level, cfg.Log.Format)

	db, err := sqlx.Connect("pgx", cfg.DSN())
	if err != nil {
		lg.Error("failed to connect to database", "error", err)
		os.Exit(1)
	} else {
		lg.Info("database connected")
	}
	defer db.Close()

	r := gin.Default()
	addr := fmt.Sprintf(":%d", cfg.Http.Port)

	r.GET("/handler", handler.Health())

	lg.Info("server started", "addr", addr)

	if err := r.Run(addr); err != nil {
		lg.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
