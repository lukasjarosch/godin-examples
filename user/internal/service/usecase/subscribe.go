package usecase

import (
	"context"
)

type subscriptionHandler struct {
	logger Logger
}

func NewSubscriptionHandler(logger Logger) *subscriptionHandler {
	return &subscriptionHandler{
		logger: logger,
	}
}

func (pub *subscriptionHandler) UserCreated(ctx context.Context, user *User) error {
	pub.logger.Info("received user.created event", "user.name", user.Username, "user.id", user.ID, "user.email", user.Email)
	return nil
}
