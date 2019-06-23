package user

import (
	"context"
	"github.com/lukasjarosch/godin-examples/user/internal/service"
	"github.com/lukasjarosch/godin/pkg/log"
	"github.com/rs/xid"
	"fmt"
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
func (s *serviceImplementation) Create(ctx context.Context, username string, email string) (user *service.UserEntity, err error) {
	user = &service.UserEntity{
		ID: xid.New().String(),
		Name: username,
		Email: email,
	}
	return user, nil
}

// Delete a user (soft-delete)
func (s *serviceImplementation) Delete(ctx context.Context, id string) (err error) {
	return fmt.Errorf("NOT_IMPLEMENTED")
}

// Get will return a user given it's ID if it exists.
func (s *serviceImplementation) Get(ctx context.Context, id string) (user *service.UserEntity, err error) {
	return nil, fmt.Errorf("NOT_IMPLEMENTED")
}

// List all registered users
func (s *serviceImplementation) List(ctx context.Context) (users []*service.UserEntity, err error) {
	return nil, fmt.Errorf("NOT_IMPLEMENTED")
}
