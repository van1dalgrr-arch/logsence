package service

import (
	"context"
	"time"

	"logsence/internal/domain"
	"logsence/internal/repository"
)

type LogService struct {
	repo repository.Logsence
}

func NewLogService(repo repository.Logsence) *LogService {
	return &LogService{
		repo: repo,
	}
}

func (s *LogService) Create(ctx context.Context, l *domain.Log) error {
	if !l.Level.Valid() {
		return domain.ErrInvalidLevel
	}

	if l.Message == "" {
		return domain.ErrEmptyField
	}

	if l.Service == "" {
		return domain.ErrEmptyField
	}

	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	return s.repo.Create(ctx, l)
}
