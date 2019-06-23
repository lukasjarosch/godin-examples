package middleware

import (
	"context"
	"time"

	"github.com/lukasjarosch/godin/pkg/log"

	"github.com/lukasjarosch/godin-examples/user/internal/service"
	"github.com/lukasjarosch/godin-examples/user/internal/service/endpoint"
)

type loggingMiddleware struct {
	next   service.User
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.User) service.User {
		return &loggingMiddleware{next, logger}
	}
}

// Create logs the request and response of the service.Create endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Create(ctx context.Context, username string, email string) (user *service.UserEntity, err error) {
	l.logger.Log(
		"endpoint", "Create",
		"request", endpoint.CreateRequest{
			Username: username,
			Email:    email,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.CreateResponse{User: user}

		l.logger.Log(
			"endpoint", "Create",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(ctx, username, email)
}

// Get logs the request and response of the service.Get endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Get(ctx context.Context, id string) (user *service.UserEntity, err error) {
	l.logger.Log(
		"endpoint", "Get",
		"request", endpoint.GetRequest{
			Id: id,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.GetResponse{User: user}

		l.logger.Log(
			"endpoint", "Get",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Get(ctx, id)
}

// List logs the request and response of the service.List endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) List(ctx context.Context) (users []*service.UserEntity, err error) {
	l.logger.Log(
		"endpoint", "List",
		"request", endpoint.ListRequest{},
	)

	defer func(begin time.Time) {
		resp := endpoint.ListResponse{Users: users}

		l.logger.Log(
			"endpoint", "List",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.List(ctx)
}

// Delete logs the request and response of the service.Delete endpoint
// The runtime will also be logged. Once a request enters this middleware, the timer is started.
// Upon leaving this middleware (deferred function is called), the time-delta is calculated.
func (l loggingMiddleware) Delete(ctx context.Context, id string) (err error) {
	l.logger.Log(
		"endpoint", "Delete",
		"request", endpoint.DeleteRequest{
			Id: id,
		},
	)

	defer func(begin time.Time) {
		resp := endpoint.DeleteResponse{}

		l.logger.Log(
			"endpoint", "Delete",
			"response", resp,
			"error", err,
			"success", err == nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Delete(ctx, id)
}
