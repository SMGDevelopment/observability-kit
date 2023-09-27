package metrics

import (
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

func MetricError(errMsg string) {
	red.Errors.WithLabelValues(errMsg).Inc()
}
