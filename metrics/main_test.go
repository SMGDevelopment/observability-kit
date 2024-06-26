package metrics

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const metricsPrefix = "rmp_observability_kit"

var metrics Metrics

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
	metrics = InitMetrics(metricsPrefix)
}

func prepHTTPCall() (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/test", nil)
	return w, r
}
