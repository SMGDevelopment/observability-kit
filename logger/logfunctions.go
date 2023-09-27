package logger

import "context"

func LogErrorContext(ctx context.Context, msg string, args ...any) {
	logger.ErrorContext(ctx, msg, args...)
}

func LogError(msg string, args ...any) {
	logger.Error(msg, args...)
}
