package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
		},
		[]string{"serivce", "path", "method", "status"},
	)

	HTTPRequestsDurationHistorgram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_historgram",
			Buckets: []float64{
				0.1,  // 100 ms
				0.2,  // 200 ms
				0.25, // 250 ms
				0.5,  // 500 ms
				1,    // 1 s
			},
		},
		[]string{"serivce", "path", "method"},
	)

	GrpcRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
		},
		[]string{"serivce", "procedure", "status"}, // status = OK | FAIL
	)

	GrpcRequestsDurationHistorgram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "grpc_request_duration_seconds_historgram",
			Buckets: []float64{
				0.1,  // 100 ms
				0.2,  // 200 ms
				0.25, // 250 ms
				0.5,  // 500 ms
				1,    // 1 s
			},
		},
		[]string{"serivce", "procedure"},
	)
)
