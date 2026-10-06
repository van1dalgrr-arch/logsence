package handler

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"logsence/internal/domain"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	service domain.LogService
	log     *slog.Logger
}

type createLogRequest struct {
	Service   string       `json:"service"`
	Level     domain.Level `json:"level"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at"`
}

// Constructor
func NewLogHandler(service domain.LogService, log *slog.Logger) *LogHandler {
	return &LogHandler{
		service: service,
		log:     log,
	}
}

type Ping interface {
	PingContext(ctx context.Context) error
}

func Health(db Ping) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.PingContext(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}
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
			h.writeError(c, err)
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
			h.writeError(c, err)
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
			h.writeError(c, err)
			return
		}
		c.JSON(200, logs)
	}
}

func (h *LogHandler) writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidLevel), errors.Is(err, domain.ErrEmptyField):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		h.log.Error(
			"request failed",
			"method", c.Request.Method,
			"path", c.FullPath(),
			"error", err,
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
