package metrics

import (
	"net/http"
	"time"

	"github.com/rabellamy/promstrap/strategy"
)

var red strategy.RED

func InitMetrics(nameSpace string) {
	//define metrics values
	redTemp, err := strategy.NewRED(strategy.REDOpts{
		RequestType:    "http",
		Namespace:      nameSpace,
		RequestLabels:  []string{"path", "verb"},
		DurationLabels: []string{"path"},
	})
	if err != nil {
		panic(err)
	}

	// register metrics
	if err := redTemp.Register(); err != nil {
		panic(err)
	}

	red = *redTemp
}

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

func MetricError(errMsg string) {
	red.Errors.WithLabelValues(errMsg).Inc()
}
