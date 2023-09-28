package logger

import (
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
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

func TestAttr(t *testing.T) {
	attr := Attr("test", "value")
	unwrappedAttr := unwrapAttr(attr)
	require.Equal(t, unwrappedAttr.Key, "test")
	require.Equal(t, unwrappedAttr.Value.String(), "value")
}

func TestUnwrapAttrs(t *testing.T) {
	attrOne := Attr("test1", "value1")
	attrTwo := Attr("test2", "value2")

	attrs := unwrapAttrs(attrOne, attrTwo)
	require.True(t, len(attrs) == 2)

	attrsOne, ok := attrs[0].(slog.Attr)
	require.True(t, ok)
	attrsTwo, ok := attrs[1].(slog.Attr)
	require.True(t, ok)

	require.Equal(t, attrsOne.Key, attrOne.key)
	require.Equal(t, attrsOne.Value.String(), attrOne.value)
	require.Equal(t, attrsTwo.Key, attrTwo.key)
	require.Equal(t, attrsTwo.Value.String(), attrTwo.value)
}
