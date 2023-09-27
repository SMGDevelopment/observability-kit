package logger

import (
	"context"
	"io"
	"log/slog"
	"strings"
)

type contextKey string

const ctxLogKey contextKey = "logFields"
const (
	EnvProd    = "PROD"
	EnvStaging = "STAGING"
	EnvDev     = "DEV"
)

type ctxLogVal map[string]any

type Logger struct {
	log *slog.Logger
}

func InitLogger(envLevel string, w io.Writer) Logger {
	opts := slog.HandlerOptions{
		AddSource: true,
	}

	// no default required as log level is by default Info
	switch strings.ToLower(envLevel) {
	case strings.ToLower(EnvDev):
		opts.Level = slog.LevelDebug
	case strings.ToLower(EnvStaging): // eventually Dev environment would be level info too
		opts.Level = slog.LevelInfo
	case strings.ToLower(EnvProd):
		opts.Level = slog.LevelError
	}

	handler := sHandler{handler: slog.NewJSONHandler(w, &opts)}
	return Logger{log: slog.New(&handler)}
}

/********* DEBUG *********/

func (logger Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	logger.log.DebugContext(ctx, msg, args...)
}

func (logger Logger) Debug(msg string, args ...any) {
	logger.log.Debug(msg, args...)
}

/********* INFO *********/

func (logger Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	logger.log.InfoContext(ctx, msg, args...)
}

func (logger Logger) Info(msg string, args ...any) {
	logger.log.Info(msg, args...)
}

/********* WARN *********/

func (logger Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	logger.log.WarnContext(ctx, msg, args...)
}

func (logger Logger) Warn(msg string, args ...any) {
	logger.log.Warn(msg, args...)
}

/********* ERROR *********/

func (logger Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	logger.log.ErrorContext(ctx, msg, args...)
}

func (logger Logger) Error(msg string, args ...any) {
	logger.log.Error(msg, args...)
}
