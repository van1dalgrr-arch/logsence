package handler

import (
	"net/http"
	"strconv"
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

func (h *LogHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		l, err := h.service.GetByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, l)
	}
}

func (h *LogHandler) ListByService() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := c.Query("service")

		limit := 0
		if limitStr := c.Query("limit"); limitStr != "" {
			var err error
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid limit"})
				return
			}
		}

		logs, err := h.service.ListByService(c.Request.Context(), service, limit)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, logs)
	}
}
