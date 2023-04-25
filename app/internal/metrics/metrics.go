package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	InFlightRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "telegram",
		Subsystem: "http",
		Name:      "in_flight_requests_total",
	})
	SummaryResponseTime = promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "telegram",
		Subsystem: "http",
		Name:      "summary_response_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	HistogramResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "telegram",
			Subsystem: "http",
			Name:      "histogram_response_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
		},
		[]string{"code"},
	)
)
