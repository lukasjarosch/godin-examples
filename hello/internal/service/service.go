package service

import (
	"context"
)

// Hello documentation is automatically
// added to the README.
type Hello interface {
	// Hello greets you. This comment is also automatically added to the README.
	// Also make sure that all parameters are named, Godin requires this information in order to work.
	Hello(ctx context.Context, name string) (greeting string, err error)
}
