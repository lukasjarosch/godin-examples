package subscriber

import (
	"context"
	"github.com/go-godin/log"
	"github.com/go-godin/rabbitmq"
	"github.com/lukasjarosch/godin-examples/hello/internal/service"
)

func UserCreatedHandler(logger log.Logger, usecase service.Hello) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		logger = logger.With("requestId", ctx.Value("requestId"))
		logger.Info("calling usecase.Hello")

		// TODO: unmarshal
		// delivery.IncrementTransportErrorCounter("user.created")

		greeting, err := usecase.Hello(ctx, "derp")
		if err != nil {
		    logger.Error("failed to call usecase.Hello()", "err", err)
			delivery.NackDelivery(false, false, "user.created")
		    return
		}

		logger.Info(greeting)
		delivery.Ack(false)
	}
}
