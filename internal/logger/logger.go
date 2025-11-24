package logger

import "context"

// Logger defines the logging interface for the application
type Logger interface {
	Info(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Debug(ctx context.Context, msg string, args ...any)
}
