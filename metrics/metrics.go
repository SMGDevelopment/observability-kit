package metrics

import (
	"github.com/rabellamy/promstrap/strategy"
)

type Metrics struct {
	red *strategy.RED
}

func InitMetrics(nameSpace string) Metrics {
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

	return Metrics{
		red: redTemp,
	}
}

func (metrics Metrics) ErrorInc(errMsg string) {
	metrics.red.Errors.WithLabelValues(errMsg).Inc()
}
