package middleware

import (
	"context"
	"time"

	grpc_metadata "github.com/go-godin/grpc-metadata"

	"github.com/go-godin/log"
	"github.com/go-kit/kit/endpoint"
)

func Logging(logger log.Logger, endpointName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Info("endoint finished",
					"endpoint", endpointName,
					"took", time.Since(begin),
					"requestId", grpc_metadata.GetRequestID(ctx),
					"transport", "grpc")
			}(time.Now())

			return next(ctx, request)
		}
	}
}
