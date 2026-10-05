package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	lg := logger.New(cfg.Log.Level, cfg.Log.Format)

	db, err := sqlx.Connect("pgx", cfg.DSN())
	if err != nil {
		lg.Error("failed to connect to database", "error", err)
		os.Exit(1)
	} else {
		lg.Info("database connected")
	}
	defer func() {
		if err := db.Close(); err != nil {
			lg.Error("failed to close database", "error", err)
		}
	}()

	logRepo := repository.NewLogRepo(db)
	logService := service.NewLogService(logRepo)
	logHandler := handler.NewLogHandler(logService, lg)

	r := gin.Default()
	addr := fmt.Sprintf(":%d", cfg.Http.Port)

	r.GET("/health", handler.Health())
	r.POST("/logs", logHandler.Create())
	r.GET("/logs/:id", logHandler.GetByID())
	r.GET("/logs", logHandler.ListByService())

	lg.Info("server started", "addr", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lg.Info("shutting down")
	if err := srv.Shutdown(ctx); err != nil {
		lg.Error("graceful shutdown failed", "error", err)
	}
	lg.Info("server stopped")
}
