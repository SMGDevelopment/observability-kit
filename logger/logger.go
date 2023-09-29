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

type LogAttr struct {
	key   string
	value any
}

func InitLogger(envLevel string, w io.Writer) Logger {
	opts := slog.HandlerOptions{}

	// no default required as log level is by default Info
	switch strings.ToLower(envLevel) {
	case strings.ToLower(EnvDev):
		opts.Level = slog.LevelDebug
	case strings.ToLower(EnvStaging): // eventually Dev environment would be level info too
		opts.Level = slog.LevelInfo
	case strings.ToLower(EnvProd):
		opts.Level = slog.LevelWarn
	}

	handler := sHandler{handler: slog.NewJSONHandler(w, &opts)}
	return Logger{log: slog.New(&handler)}
}

func Attr(key string, value any) LogAttr {
	return LogAttr{key: key, value: value}
}

func unwrapAttr(attr LogAttr) slog.Attr {
	return slog.Any(attr.key, attr.value)
}

func unwrapAttrs(logAttrs ...LogAttr) []any {
	attrs := make([]any, 0)
	for _, v := range logAttrs {
		attrs = append(attrs, unwrapAttr(v))
	}
	return attrs
}

/********* DEBUG *********/

func (logger Logger) DebugContext(ctx context.Context, msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.DebugContext(ctx, msg, attrs...)
}

func (logger Logger) Debug(msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.Debug(msg, attrs...)
}

/********* INFO *********/

func (logger Logger) InfoContext(ctx context.Context, msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.InfoContext(ctx, msg, attrs...)
}

func (logger Logger) Info(msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.Info(msg, attrs...)
}

/********* WARN *********/

func (logger Logger) WarnContext(ctx context.Context, msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.WarnContext(ctx, msg, attrs...)
}

func (logger Logger) Warn(msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.Warn(msg, attrs...)
}

/********* ERROR *********/

func (logger Logger) ErrorContext(ctx context.Context, msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.ErrorContext(ctx, msg, attrs...)
}

func (logger Logger) Error(msg string, args ...LogAttr) {
	attrs := unwrapAttrs(args...)
	logger.log.Error(msg, attrs...)
}
