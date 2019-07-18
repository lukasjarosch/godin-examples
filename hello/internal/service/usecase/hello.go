package usecase

import (
	"context"
	"fmt"

	"github.com/go-godin/log"
	_ "github.com/lukasjarosch/godin-examples/hello/internal/service"
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
	return fmt.Sprintf("Ohai there, %s", name), nil
}
