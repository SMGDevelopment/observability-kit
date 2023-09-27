package logger

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name     string
		envLevel string
	}{
		{
			name:     "DEV",
			envLevel: "DEV",
		},
		{
			name:     "STAGING",
			envLevel: "STAGING",
		},
		{
			name:     "PROD",
			envLevel: "PROD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger(tt.envLevel, os.Stdout)
		})
	}
}

func TestLogRequestIDMiddleware(t *testing.T) {
	// create a handler to use as "next" which will verify the request
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// create the handler to test, using our custom "next" handler
	handlerToTest := LogRequestIDMiddleware(nextHandler)

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(prepHTTPCall())
}

func TestLogDebug(t *testing.T) {
	logger.DebugContext(context.TODO(), "my error message")
	logger.Debug("my error message")
}

func TestLogInfo(t *testing.T) {
	logger.InfoContext(context.TODO(), "my error message")
	logger.Info("my error message")
}

func TestLogWarn(t *testing.T) {
	logger.WarnContext(context.TODO(), "my error message")
	logger.Warn("my error message")
}

func TestLogError(t *testing.T) {
	logger.ErrorContext(context.TODO(), "my error message")
	logger.Error("my error message")
}
