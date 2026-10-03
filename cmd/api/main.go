package main

import (
	"fmt"
	"os"

	"logsence/internal/config"
	"logsence/internal/handler"
	"logsence/internal/repository"
	"logsence/internal/service"
	"logsence/pkg/logger"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		panic(fmt.Errorf("failed to load conifg: %w", err))
	}

	lg := logger.New(cfg.Log.Level, cfg.Log.Format)

	db, err := sqlx.Connect("pgx", cfg.DSN())
	if err != nil {
		lg.Error("failed to connect to database", "error", err)
		os.Exit(1)
	} else {
		lg.Info("database connected")
	}
	defer db.Close()

	logRepo := repository.NewLogRepo(db)
	logService := service.NewLogService(logRepo)
	logHandler := handler.NewLogHandler(logService, lg)

	r := gin.Default()
	addr := fmt.Sprintf(":%d", cfg.Http.Port)

	r.GET("/handler", handler.Health())
	r.POST("/logs", logHandler.Create())
	r.GET("/logs/:id", logHandler.GetByID())
	r.GET("/logs", logHandler.ListByService())

	lg.Info("server started", "addr", addr)

	if err := r.Run(addr); err != nil {
		lg.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
