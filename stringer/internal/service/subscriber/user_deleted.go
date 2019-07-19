package subscriber

import (
	"context"
	grpcMetadata "github.com/go-godin/grpc-metadata"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	"github.com/pkg/errors"

	userDeletedProto "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

// UserDeleted is responsible of handling all incoming AMQP messages with routing key 'user.deleted'
// It might seem overly complicated at first, but the design is on purpose. You WANT to have access to the Delivery,
// thus it would not make sense to use a middleware for Decoding it into a DAO or domain-level object as you would
// loose access to the Delivery.
func UserDeletedSubscriber(logger log.Logger, usecase service.Stringer, decoder rabbitmq.SubscriberDecoder) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger = logger.With(string(grpcMetadata.RequestID), ctx.Value(string(grpcMetadata.RequestID)))

		event, err := decodeUserDeleted(delivery, decoder, logger)
		if err != nil {
			return
		}

		_ = event // remove this line, it just keeps the compiler calm until you start using the event :)

		// TODO: Handle user.deleted subscription here
		/*
			If you want to NACK the delivery, use `delivery.NackDelivery()` instead of Nack().
			This will ensure that the prometheus amqp_nack_counter is increased.
		*/

		_ = delivery.Ack(false)
	}
}

// decodeUserDeleted cleans up the actual handler by providing a cleaner interface for decoding incoming UserCreatedEvent deliveries.
// It will also take care of logging errors and handling metrics.
func decodeUserDeleted(delivery *rabbitmq.Delivery, decoder rabbitmq.SubscriberDecoder, logger log.Logger) (*userDeletedProto.UserCreatedEvent, error) {
	event, err := decoder(delivery)
	if err != nil {
		if err2 := delivery.NackDelivery(false, false, "user.deleted"); err2 != nil {
			err = errors.Wrap(err, err2.Error())
		}
		delivery.IncrementTransportErrorCounter("user.deleted")
		logger.Error("failed to decode UserCreatedEvent", "err", err)
		return nil, err
	}
	logger.Debug("decoded UserCreatedEvent", "event", event)

	return event.(*userDeletedProto.UserCreatedEvent), nil
}
