package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidLevel = errors.New("invalid log level")
	ErrNotFound = errors.New("log not found")
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
)

func (l Level) Valid() bool {
	switch l {
	case LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal:
		return true
	}
	return false
}

type Log struct {
	ID         int64     `json:"id"          db:"id"`
	Service    string    `json:"service"     db:"service"`
	Level      Level     `json:"level"       db:"level"`
	Message    string    `json:"message"     db:"message"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
	ReceivedAt time.Time `json:"received_at" db:"received_at"`
}
