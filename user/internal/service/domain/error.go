package domain

import "errors"

// Domain errors
// These can then be remaped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
	ErrNotImplemented = errors.New("use-case not implemented")

	EmailEmptyError    = errors.New("the email address may not be empty")
	ErrSomethingFailed = errors.New("something failed")

	ErrMalformedEvent = errors.New("malformed event, cannot umarshal")
)
