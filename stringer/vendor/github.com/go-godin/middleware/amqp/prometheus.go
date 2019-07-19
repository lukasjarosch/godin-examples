package amqp

import (
	"context"

	"github.com/go-godin/rabbitmq"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var amqpInbound = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
	Name: "amqp_inbound",
	Help: "Increased on incoming message, decreased on ack/nack",
}, []string{"routing_key"})

func PrometheusInstrumentation(routingKey string, handler rabbitmq.SubscriptionHandler) rabbitmq.SubscriptionHandler {
	return func(ctx context.Context, delivery *rabbitmq.Delivery) {
		amqpInbound.With("routing_key", routingKey).Add(1)

		defer func() {
			amqpInbound.With("routing_key", routingKey).Add(-1)
		}()

		handler(ctx, delivery)
	}
}
