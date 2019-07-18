package middleware

import (
	"context"
	"time"

	"github.com/go-godin/log"

	"github.com/lukasjarosch/godin-examples/greeter2/internal/service"
	"github.com/lukasjarosch/godin-examples/greeter2/internal/service/endpoint"
)

type loggingMiddleware struct {
	next   service.Greeter
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Greeter) service.Greeter {
		return &loggingMiddleware{next, logger}
	}
}

// Hello logs the request and response of the service.Hello endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Hello(ctx context.Context, name string) (greeting string, err error) {
	l.logger.Log(
		"endpoint", "Hello",
		"request", endpoint.HelloRequest{
			Name: name,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.HelloResponse{Greeting: greeting}

		l.logger.Log(
			"endpoint", "Hello",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Hello(ctx, name)
}
