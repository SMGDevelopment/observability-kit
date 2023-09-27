package logger

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var logger Logger

func TestMain(m *testing.M) {
	logger = InitLogger("DEV", os.Stdout)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func prepHTTPCall() (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/test", nil)
	return w, r
}
