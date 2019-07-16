package usecase

import (
	"context"
)

// Hello greets you. This comment is also automatically added to the README.
// Also make sure that all parameters are named, Godin requires this information in order to work.
func (s *greeterInteractor) Hello(ctx context.Context, name string) (greeting Greeting, err error) {
	s.Hello()
}
