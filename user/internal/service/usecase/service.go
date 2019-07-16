package usecase

import (
	"context"
	"fmt"
	"github.com/rs/xid"

	"github.com/lukasjarosch/godin-examples/user/internal/service/domain"
)

// Service documentation is automatically
// added to the README.
type Service interface {
	// Create will create a new user and return it.
	Create(ctx context.Context, username string, email string) (user *User, err error)
	// Get will return a user given it's ID if it exists.
	Get(ctx context.Context, id string) (user *User, err error)
	// List all registered users
	List(ctx context.Context) (users []*User, err error)
	// Delete a user (soft-delete)
	Delete(ctx context.Context, id string) (err error)
}

type Publisher interface {
	UserCreated(ctx context.Context, user *User) error
}

type Subscriber interface {
	UserCreated(ctx context.Context, user *User) error
}

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type User struct {
	ID       string
	UserID   string
	Username string
	Email    string
}

func toUser(user *domain.User) *User {
	return &User{
		ID:       xid.New().String(),
		Username: "asdf",
		Email:    "email",
		UserID:   fmt.Sprintf("user_%s", xid.New().String()),
	}
}
