package middleware

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"time"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc/status"
)

func PrometheusMiddleware(duration metrics.Histogram, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(now time.Time) {
				code, _ := status.FromError(err)
				duration.With("method", methodName, "status_code", code.Code().String())
			}(time.Now())
			return next(ctx, request)
		}
	}
}
