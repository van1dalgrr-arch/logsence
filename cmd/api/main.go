package main

import (
	"fmt"
	"os"

	"logsence/internal/config"
	"logsence/internal/handler"
	"logsence/pkg/loger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load("config.yml")

	if err != nil {
		panic(fmt.Errorf("failed to load conifg: %w", err))
	}

	lg := loger.New(cfg.Log.Level, cfg.Log.Format)

	r := gin.Default()
	addr := fmt.Sprintf(":%d", cfg.Http.Port)

	r.GET("/handler", handler.Health())

	lg.Info("server started", "addr", addr)

	if err := r.Run(addr); err != nil {
		lg.Error("server failed to start", "error", err)
		os.Exit(1)
	}

}
