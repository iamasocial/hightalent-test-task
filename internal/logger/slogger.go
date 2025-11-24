package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// Slogger is a slog-based implementation of Logger interface
type Slogger struct {
	logger *slog.Logger
}

// NewSlogger creates a slog-based Logger instance.
// The env parameter controls the log format and level
// - "local": human-readable text, debug level
// - "dev": JSON format, debug level
// - "prod": JSON format, info level
func NewSlogger(env string) Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	log.With(slog.String("env", env))

	return &Slogger{
		logger: log,
	}
}

// Info logs at LevelInfo
func (s *Slogger) Info(ctx context.Context, msg string, args ...any) {
	s.logger.Info(msg, args...)
}

// Error logs at LevelError
func (s *Slogger) Error(ctx context.Context, msg string, args ...any) {
	s.logger.Error(msg, args...)
}

// Debug logs at LevelDebug
func (s *Slogger) Debug(ctx context.Context, msg string, args ...any) {
	s.logger.Debug(msg, args...)
}
