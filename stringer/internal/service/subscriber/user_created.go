package subscriber

import (
	"context"
	grpcMetadata "github.com/go-godin/grpc-metadata"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	"github.com/pkg/errors"

	userCreatedProto "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

// UserCreated is responsible of handling all incoming AMQP messages with routing key 'user.created'
func UserCreatedSubscriber(logger log.Logger, usecase service.Stringer, decoder rabbitmq.SubscriberDecoder) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger = logger.With(string(grpcMetadata.RequestID), ctx.Value(string(grpcMetadata.RequestID)))

		event, err := decodeUserCreated(delivery, decoder, logger)
		if err != nil {
			return
		}

		_ = event // remove this line, it just keeps the compiler calm until you start using the event :)

		// TODO: Handle user.created subscription here
		/*
			If you want to NACK the delivery, use `delivery.NackDelivery()` instead of Nack().
			This will ensure that the prometheus amqp_nack_counter is increased.
		*/

		_ = delivery.Ack(false)
	}
}

// decodeUserCreated cleans up the actual handler by providing a cleaner interface for decoding incoming UserCreatedEvent deliveries.
// It will also take care of logging errors and handling metrics.
func decodeUserCreated(delivery *rabbitmq.Delivery, decoder rabbitmq.SubscriberDecoder, logger log.Logger) (*userCreatedProto.UserCreatedEvent, error) {
	event, err := decoder(delivery)
	if err != nil {
		if err2 := delivery.NackDelivery(false, false, "user.created"); err2 != nil {
			err = errors.Wrap(err, err2.Error())
		}
		delivery.IncrementTransportErrorCounter("user.created")
		logger.Error("failed to decode UserCreatedEvent", "err", err)
		return nil, err
	}
	logger.Debug("decoded UserCreatedEvent", "event", event)

	return event.(*userCreatedProto.UserCreatedEvent), nil
}
