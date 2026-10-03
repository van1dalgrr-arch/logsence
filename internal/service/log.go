package service

import (
	"context"
	"time"

	"logsence/internal/domain"
	"logsence/internal/repository"
)

var _ domain.LogService = (*LogService)(nil)

const (
	defaultLimit = 10
	maxLimit     = 1000
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

func (s *LogService) GetByID(ctx context.Context, id int64) (*domain.Log, error) {
	if id <= 0 {
		return nil, domain.ErrNotFound
	}
	return s.repo.GetByID(ctx, id)
}

func (s *LogService) ListByService(ctx context.Context, service string, limit int) ([]domain.Log, error) {
	if service == "" {
		return nil, domain.ErrEmptyField
	}
	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	return s.repo.ListByService(ctx, service, limit)
}
