package amqp

import (
	"context"
	"fmt"

	grpc_metadata "github.com/go-godin/grpc-metadata"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
)

// Logging provides a simple logging middleware for AMQP deliveries. It should be registered AFTER the RequestID middleware
// in order to log the requestId value properly.
func Logging(logger log.Logger, routingKey string, handler rabbitmq.SubscriptionHandler) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger.Info(
			"consume message",
			"routing_key", routingKey,
			"redelivered", fmt.Sprint(delivery.Redelivered),
			"requestId", grpc_metadata.GetRequestID(ctx),
			"transport", "amqp",
		)

		handler(ctx, delivery)
	}
}
