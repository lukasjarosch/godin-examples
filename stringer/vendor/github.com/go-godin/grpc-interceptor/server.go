package grpc_interceptor

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_ms",
			Help:    "Duration of gRPC requests in ms",
			Buckets: []float64{50, 100, 250, 1000},
		},
		[]string{"method", "status_code"})
)

func init() {
	prometheus.Register(requestDuration)
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	h, err := handler(ctx, req)

	defer func(begin time.Time) {
		stat, _ := status.FromError(err)

		meth := strings.Split(info.FullMethod, "/")

		requestDuration.With(prometheus.Labels{
			"method":      meth[len(meth)-1],
			"status_code": stat.Code().String(),
		}).Observe(time.Since(begin).Seconds() / 1000)

	}(start)

	return h, err
}
