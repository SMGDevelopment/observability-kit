package metrics

import (
	"net/http"
	"time"

	"github.com/rabellamy/promstrap/strategy"
)

var RED strategy.RED

func InitMetrics(nameSpace string) {
	//define metrics values
	red, err := strategy.NewRED(strategy.REDOpts{
		RequestType:    "http",
		Namespace:      nameSpace,
		RequestLabels:  []string{"path", "verb"},
		DurationLabels: []string{"path"},
	})
	if err != nil {
		panic(err)
	}

	// register metrics
	if err := red.Register(); err != nil {
		panic(err)
	}

	RED = *red
}

func REDMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		path := r.URL.Path

		// Records that a request took place
		RED.Requests.WithLabelValues(path, "GET").Inc()

		next.ServeHTTP(w, r)

		// Records the duration of the request
		RED.Duration.Histogram.WithLabelValues(path).Observe(time.Since(t).Seconds())
	})
}
