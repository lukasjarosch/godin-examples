package user

import (
	"context"
	"fmt"

	. "github.com/lukasjarosch/godin-examples/user/internal/service"
	"github.com/lukasjarosch/godin/pkg/log"
	"github.com/rs/xid"
)

type serviceImplementation struct {
	logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
	return &serviceImplementation{
		logger: logger,
	}
}

// Create will create a new user and return it.
func (s *serviceImplementation) Create(ctx context.Context, username string, email string) (user *UserEntity, err error) {
	user = &UserEntity{
		ID:    xid.New().String(),
		Name:  username,
		Email: email,
	}
	return user, nil
}

// Delete a user (soft-delete)
func (s *serviceImplementation) Delete(ctx context.Context, id string) (err error) {
	return fmt.Errorf("NOT_IMPLEMENTED")
}

// Get will return a user given it's ID if it exists.
func (s *serviceImplementation) Get(ctx context.Context, id string) (user *UserEntity, err error) {
	return nil, ErrNotImplemented
}

// List all registered users
func (s *serviceImplementation) List(ctx context.Context) (users []*UserEntity, err error) {
	return nil, fmt.Errorf("NOT_IMPLEMENTED")
}
