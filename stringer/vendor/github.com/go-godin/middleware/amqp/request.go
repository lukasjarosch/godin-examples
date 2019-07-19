package amqp

import (
	"context"

	grpc_metadata "github.com/go-godin/grpc-metadata"
	"github.com/go-godin/rabbitmq"
	"github.com/google/uuid"
)

func RequestID(handler rabbitmq.SubscriptionHandler) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		reqId := delivery.Headers[string(grpc_metadata.RequestID)]

		if reqId == nil {
			reqId = uuid.New().String()
		}

		ctx = context.WithValue(ctx, string(grpc_metadata.RequestID), reqId)

		handler(ctx, delivery)
	}
}
