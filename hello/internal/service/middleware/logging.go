package middleware

import (
	"context"
	"time"

	grpc_metadata "github.com/go-godin/grpc-metadata"

	"github.com/go-godin/log"

	"github.com/lukasjarosch/godin-examples/hello/internal/service"
	"github.com/lukasjarosch/godin-examples/hello/internal/service/endpoint"
)

type loggingMiddleware struct {
	next   service.Hello
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Hello) service.Hello {
		return &loggingMiddleware{next, logger}
	}
}

// Hello logs the request and response of the service.Hello endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Hello(ctx context.Context, name string) (greeting string, err error) {
	l.logger.Log(
		"endpoint", "Hello",
		"request.id", grpc_metadata.GetRequestID(ctx),
		"request.data", endpoint.HelloRequest{
			Name: name,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.HelloResponse{Greeting: greeting}

		l.logger.Log(
			"endpoint", "Hello",
			"request.id", grpc_metadata.GetRequestID(ctx),
			"response.data", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Hello(ctx, name)
}
