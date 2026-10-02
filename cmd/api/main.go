package main

import (
	"fmt"
	"log"

	"logsence/internal/config"
	"logsence/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load("config.yml")

	if err != nil {
		panic(fmt.Errorf("failed to load conifg: %w", err))
	}

	r := gin.Default()
	addr := fmt.Sprintf(":%d", cfg.Http.Port)

	r.GET("/handler", handler.Health())

	if err := r.Run(addr); err != nil {
		log.Fatal("server failed to stop", err)
	}

}
