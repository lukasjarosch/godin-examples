package yyy

import (
	"context"
	"fmt"

	"github.com/lukasjarosch/godin/pkg/log"
)

type serviceImplementation struct {
	logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
	return &serviceImplementation{
		logger: logger,
	}
}

// Hello greets you. This comment is also automatically added to the README.
// Also make sure that all parameters are named, Godin requires this information in order to work.
func (s *serviceImplementation) Hello(ctx context.Context, name string) (greeting string, err error) {
	return "", fmt.Errorf("NOT_IMPLEMENTED")
}
