package usecase

import (
	"context"
	grpc_metadata "github.com/go-godin/grpc-metadata"
	"github.com/lukasjarosch/godin-examples/stringer/internal/amqp"
	user "github.com/lukasjarosch/godin-examples/user/api"

	"github.com/lukasjarosch/godin-examples/stringer/internal/service/domain"

	"github.com/go-godin/log"
	_ "github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

type serviceImplementation struct {
	logger     log.Logger
	publishers amqp.PublisherSet
}

func NewServiceImplementation(logger log.Logger, publishers amqp.PublisherSet) *serviceImplementation {
	return &serviceImplementation{
		logger:     logger,
		publishers: publishers,
	}
}

// Hello greets you. This comment is also automatically added to the README.
// Also make sure that all parameters are named, Godin requires this information in order to work.
func (s *serviceImplementation) Hello(ctx context.Context, name string) (greeting string, err error) {
	//return fmt.Sprintf("ohai there, %s", name), nil
	if err := s.publishers.PublishUserCreated(ctx, &user.UserCreated{}); err != nil {

	}
	/*
		if err := s.userCreatedPublisher.Publish(ctx, &user.UserCreated{}); err != nil {
			s.logger.Error("failed to publish event", "err", err)
		}
	*/
	s.logger.Info("published event", "topic", "user.created", "requestId", grpc_metadata.GetRequestID(ctx))
	return "", domain.ErrUnauthenticated
}
