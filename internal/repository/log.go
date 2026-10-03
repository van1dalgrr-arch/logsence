package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"logsence/internal/domain"
)

type Logsence interface {
	Create(ctx context.Context, l *domain.Log) error
	GetByID(ctx context.Context, id int64) (*domain.Log, error)
	ListByService(ctx context.Context, service string, limit int) ([]domain.Log, error)
}

type LogsenceRepository struct {
	db *sqlx.DB
}

var _ Logsence = (*LogsenceRepository)(nil)

func NewLogRepo(db *sqlx.DB) *LogsenceRepository {
	return &LogsenceRepository{
		db: db,
	}
}

func (r *LogsenceRepository) Create(ctx context.Context, l *domain.Log) error {
	const query = `
	INSERT INTO logs (service, level, message, created_at)
	VALUES($1, $2, $3, $4)
	RETURNING id,  received_at
	`
	return r.db.QueryRowxContext(ctx, query, l.Service, l.Level, l.Message, l.CreatedAt).
		Scan(&l.ID, &l.ReceivedAt)
}

func (r *LogsenceRepository) GetByID(ctx context.Context, id int64) (*domain.Log, error) {
	const query = `
	SELECT id, service, level, message, created_at, received_at FROM logs WHERE id = $1
	`

	var l domain.Log

	err := r.db.GetContext(ctx, &l, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("repository GetByID: %w", err)
	}
	return &l, nil
}

func (r *LogsenceRepository) ListByService(ctx context.Context, service string, limit int) ([]domain.Log, error) {
	const query = `
	SELECT id, service, level, message, created_at, received_at FROM logs WHERE service = $1 ORDER BY created_at DESC LIMIT $2`

	logs := []domain.Log{}

	err := r.db.SelectContext(ctx, &logs, query, service, limit)
	if err != nil {
		return nil, fmt.Errorf("repository ListByService: %w", err)
	}
	return logs, nil
}
