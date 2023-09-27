package metrics

import (
	"net/http"
	"time"
)

func REDMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		path := r.URL.Path

		// Records that a request took place
		red.Requests.WithLabelValues(path, "GET").Inc()

		next.ServeHTTP(w, r)

		// Records the duration of the request
		red.Duration.Histogram.WithLabelValues(path).Observe(time.Since(t).Seconds())
	})
}
