package logger

import (
	"context"
	"log/slog"
)

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
