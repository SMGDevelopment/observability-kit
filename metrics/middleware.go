package metrics

import (
	"net/http"
	"time"
)

func (metrics Metrics) REDMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		path := r.URL.Path

		// Records that a request took place
		metrics.red.Requests.WithLabelValues(path, "GET").Inc()

		next.ServeHTTP(w, r)

		// Records the duration of the request
		metrics.red.Duration.Histogram.WithLabelValues(path).Observe(time.Since(t).Seconds())
	})
}
