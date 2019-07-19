package rabbitmq

import (
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	transportError = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "amqp_transport_error",
		Help: "Increased when a message could not be decoded or necessary content is missing",
	}, []string{"routing_key"})
	nackCounter = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "amqp_not_acknowledged",
		Help: "Increased by 1 on every nack on AMQP messages",
	}, []string{"routing_key", "requeue"})
)
