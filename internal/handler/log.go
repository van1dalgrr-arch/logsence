package handler

import (
	"logsence/internal/domain"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	service domain.LogService
}

// Constructor
func NewLogHandler(service domain.LogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is healthy",
		})
	}
}
