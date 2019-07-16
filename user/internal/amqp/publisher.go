package amqp

import (
	"context"

	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"

	"github.com/go-godin/log"
	kitAMQP "github.com/go-kit/kit/transport/amqp"
	"github.com/streadway/amqp"
)

var NopDecoder = func(context.Context, *amqp.Delivery) (response interface{}, err error) { return nil, nil }

type amqpPublisher struct {
	userCreated *kitAMQP.Publisher
	logger      log.Logger
}

func NewPublisher(channel *amqp.Channel, logger log.Logger) *amqpPublisher {
	options := []kitAMQP.PublisherOption{
		kitAMQP.PublisherDeliverer(kitAMQP.SendAndForgetDeliverer),
	}

	return &amqpPublisher{
		logger: logger,
		userCreated: kitAMQP.NewPublisher(
			channel,
			&amqp.Queue{},
			EncodeUserCreatedEvent,
			NopDecoder,
			append(options, kitAMQP.PublisherBefore(
				kitAMQP.SetPublishKey("user.created"),
				kitAMQP.SetPublishExchange("exchange-name"),
			))...,
		),
	}
}

func (pub *amqpPublisher) UserCreated(ctx context.Context, user *usecase.User) error {
	_, err := pub.userCreated.Endpoint()(ctx, user)
	if err != nil {
		pub.logger.Error("UserCreated subscriber endpoint failed", "err", err)
		return err
	}
	pub.logger.Info("published message to 'user.created' on exchange 'exchange-name'")
	return nil
}
