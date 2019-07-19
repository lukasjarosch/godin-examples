package middleware

import (
	"context"
	"time"

	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var requestDuration = prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
	Name:    "endpoint_request_duration_ms",
	Help:    "Request duration in milliseconds",
	Buckets: []float64{50, 100, 250, 500, 1000},
}, []string{"method"})

var requestsCurrent = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
	Name: "endpoint_requests_current",
	Help: "The current number of gRPC requests by endpoint",
}, []string{"method"})

var requestStatus = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
	Name: "endpoint_requests_total",
	Help: "The total number of gRPC requests and whether the business failed or not",
}, []string{"method", "success"})

// InstrumentGRPC adds basic RED metrics on all endpoints. The transport layer (gRPC, AMQP, HTTP, ...) should also have metrics attached and
// will then take care of monitoring gRPC endpoints including their status.
func InstrumentGRPC(methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			requestsCurrent.With("method", methodName).Add(1)

			defer func(begin time.Time) {
				requestDuration.With("method", methodName).Observe(time.Since(begin).Seconds())
				requestsCurrent.With("method", methodName).Add(-1)
				requestStatus.With("method", methodName, "success", fmt.Sprint(err == nil))
			}(time.Now())

			return next(ctx, request)
		}
	}
}
