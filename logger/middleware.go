package logger

import (
	"context"
	"log/slog"
	"net/http"

	chimw "github.com/go-chi/chi/v5/middleware"
)

// LogRequestIDMiddleware middleware for logging using requestID middleware from Chi
func LogRequestIDMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestIDKey := "requestID"
		requestID := chimw.GetReqID(r.Context())
		ctx := r.Context()
		attr := slog.String(requestIDKey, requestID)
		if ctxMap, ok := ctx.Value(ctxLogKey).(ctxLogVal); ok {
			ctxMap[requestIDKey] = attr
		} else {
			ctxLogMap := ctxLogVal{}
			ctxLogMap[requestIDKey] = attr
			ctx = context.WithValue(ctx, ctxLogKey, ctxLogMap)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return chimw.RequestID(http.HandlerFunc(fn))
}
