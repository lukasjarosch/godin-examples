package subscriber

import (
	"context"
	grpc_metadata "github.com/go-godin/grpc-metadata"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"

	userCreatedProto "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

// UserCreated is responsible of handling all incoming AMQP messages with routing key 'user.created'
func UserCreatedSubscriber(logger log.Logger, usecase service.Stringer, decoder rabbitmq.SubscriberDecoder) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		// the requestId is injected into the context and should be attached on every log
		logger = logger.With(string(grpc_metadata.RequestID), ctx.Value(string(grpc_metadata.RequestID)))

		event, err := decoder(delivery)
		event = event.(userCreatedProto.UserCreatedEvent)
		if err != nil {
			logger.Error("failed to decode 'user.created' event", "err", err)
			delivery.NackDelivery(false, false, "user.created")
			delivery.IncrementTransportErrorCounter("user.created")
			return
		}

		// TODO: Handle user.created subscription
		/*
			If you want to NACK the delivery, use `delivery.NackDelivery()` instead of Nack().
			This will ensure that the prometheus amqp_nack_counter is increased.

			Godins delivery wrapper also provides a `delivery.IncrementTransportErrorCounter()` method to grant
			you access to the amqp_transport_error metric. Call it if the message is incomplete or cannot
			be unmarshalled for any reason.
		*/

		_ = delivery.Ack(false)
	}
}
