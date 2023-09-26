package logger

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"strings"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type contextKey string

const ctxLogKey contextKey = "logFields"
const (
	EnvProd    = "PROD"
	EnvStaging = "STAGING"
	EnvDev     = "DEV"
)

type ctxLogVal map[string]any

var Logger *slog.Logger

// custom Handler implementation
type sHandler struct {
	handler slog.Handler
}

func (s *sHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return s.handler.Enabled(ctx, level)
}

func (s *sHandler) Handle(ctx context.Context, record slog.Record) error {
	if ctxMap, ok := ctx.Value(ctxLogKey).(ctxLogVal); ok {
		for _, v := range ctxMap {
			if attr, ok := v.(slog.Attr); ok {
				record.AddAttrs(attr)
			}
		}
	}
	return s.handler.Handle(ctx, record)
}

func (s *sHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return s.handler.WithAttrs(attrs)
}

func (s *sHandler) WithGroup(name string) slog.Handler {
	return s.handler.WithGroup(name)
}

func InitLogger(envLevel string, w io.Writer) {
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
	Logger = slog.New(&handler)
}

// LogRequestIDMiddleware middleware for logging using requestID middleware from Chi
func LogRequestIDMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := chimw.GetReqID(r.Context())
		ctx := r.Context()
		attr := slog.String("requestID", requestID)
		if ctxMap, ok := ctx.Value(ctxLogKey).(ctxLogVal); ok {
			ctxMap["requestID"] = attr
		} else {
			ctxLogMap := ctxLogVal{}
			ctxLogMap["requestID"] = attr
			ctx = context.WithValue(ctx, ctxLogKey, ctxLogMap)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return chimw.RequestID(http.HandlerFunc(fn))
}

func LogErrorContext(ctx context.Context, msg string, args ...any) {
	Logger.ErrorContext(ctx, msg, args...)
}

func LogError(msg string, args ...any) {
	Logger.Error(msg, args...)
}
