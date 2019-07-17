package domain

import "errors"

// Application errors
// These can then be remaped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
	ErrNotImplemented = errors.New("endpoint not implemented")
)
