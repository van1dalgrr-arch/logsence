package handler

import (
	"net/http"
	"time"

	"logsence/internal/domain"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	service domain.LogService
}

type createLogRequest struct {
	Service   string       `json:"service"`
	Level     domain.Level `json:"level"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at"`
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

func (h *LogHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createLogRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}
		l := domain.Log{
			Service:   req.Service,
			Level:     req.Level,
			Message:   req.Message,
			CreatedAt: req.CreatedAt,
		}

		if err := h.service.Create(c.Request.Context(), &l); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, l)
	}
}
