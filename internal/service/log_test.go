package service

import (
	"context"
	"errors"
	"testing"

	"logsence/internal/domain"
	"logsence/internal/repository"
)

type fakeRepo struct{}

var _ repository.Logsence = (*fakeRepo)(nil)

func (f *fakeRepo) Create(ctx context.Context, l *domain.Log) error {
	return nil
}

func (f *fakeRepo) GetByID(ctx context.Context, id int64) (*domain.Log, error) {
	return nil, nil
}

func (f *fakeRepo) ListByService(ctx context.Context, service string, limimt int) ([]domain.Log, error) {
	return nil, nil
}

func Test_EmptyMessage(t *testing.T) {
	src := NewLogService(&fakeRepo{})

	l := domain.Log{
		Service: "api",
		Message: "",
		Level:   domain.LevelInfo,
	}

	err := src.Create(t.Context(), &l)

	if !errors.Is(err, domain.ErrEmptyField) {
		t.Errorf("expected ErrEmptyField, got %v", err)
	}
}

func Test_InvalidLevel(t *testing.T) {
	src := NewLogService(&fakeRepo{})

	l := domain.Log{
		Service: "api",
		Message: "test",
		Level:   "invalid",
	}

	err := src.Create(t.Context(), &l)

	if !errors.Is(err, domain.ErrInvalidLevel) {
		t.Errorf("expected ErrInvalidLevel, got %v", err)
	}
}

func Test_ListByService(t *testing.T) {
	src := NewLogService(&fakeRepo{})

	logs, err := src.ListByService(t.Context(), "api", 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(logs) != 0 {
		t.Errorf("expected empty logs, got %v", logs)
	}
}
