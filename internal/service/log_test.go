package service

import (
	"context"
	"errors"
	"testing"

	"logsence/internal/domain"
	"logsence/internal/repository"
)

type fakeRepo struct {
	gotLimit int
}

var _ repository.Logsence = (*fakeRepo)(nil)

func (f *fakeRepo) Create(ctx context.Context, l *domain.Log) error {
	return nil
}

func (f *fakeRepo) GetByID(ctx context.Context, id int64) (*domain.Log, error) {
	return nil, nil
}

func (f *fakeRepo) ListByService(ctx context.Context, service string, limit int) ([]domain.Log, error) {
	f.gotLimit = limit
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

func Test_ListByService_DefaultLimit(t *testing.T) {
	fake := &fakeRepo{}
	src := NewLogService(fake)

	_, _ = src.ListByService(t.Context(), "api", 0)

	if fake.gotLimit != defaultLimit {
		t.Errorf("expected default limit 10, got %v", fake.gotLimit)
	}
}

func Test_ListByService_MaxLimit(t *testing.T) {
	fake := &fakeRepo{}
	src := NewLogService(fake)

	_, _ = src.ListByService(t.Context(), "api", 100000)

	if fake.gotLimit != maxLimit {
		t.Errorf("expected max limit 1000, got %v", fake.gotLimit)
	}
}
