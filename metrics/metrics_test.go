package metrics

import (
	"net/http"
	"testing"
)

func TestInitMetrics(t *testing.T) {
	InitMetrics("test")
}

func TestREDMetricsMiddleware(t *testing.T) {
	// create a handler to use as "next" which will verify the request
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// create the handler to test, using our custom "next" handler
	handlerToTest := REDMetricsMiddleware(nextHandler)

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(prepHTTPCall())
}

func TestMetricError(t *testing.T) {
	MetricError("cool error message")
}
