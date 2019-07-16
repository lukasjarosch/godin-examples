package usecase

import (
	"context"
	"fmt"

	"github.com/lukasjarosch/godin-examples/user/internal/service/domain"
)

type serviceImplementation struct {
	logger    Logger
	publisher Publisher
}

func NewServiceImplementation(logger Logger, publisher Publisher) *serviceImplementation {
	return &serviceImplementation{
		logger:    logger,
		publisher: publisher,
	}
}

// Create will create a new user and return it.
func (s *serviceImplementation) Create(ctx context.Context, username string, email string) (user *User, err error) {
	user = &User{}
	u := domain.NewUser(email, username, false)
	user = toUser(u)
	if err := s.publisher.UserCreated(ctx, user); err != nil {
		s.logger.Error("failed to publish userCreated event", "err", err)
		return nil, err
	}
	return user, nil
}

// Delete a user (soft-delete)
func (s *serviceImplementation) Delete(ctx context.Context, id string) (err error) {
	return fmt.Errorf("NOT_IMPLEMENTED")
}

// Get will return a user given it's ID if it exists.
func (s *serviceImplementation) Get(ctx context.Context, id string) (user *User, err error) {
	return nil, domain.ErrNotImplemented
}

// List all registered users
func (s *serviceImplementation) List(ctx context.Context) (users []*User, err error) {
	return nil, fmt.Errorf("NOT_IMPLEMENTED")
}
