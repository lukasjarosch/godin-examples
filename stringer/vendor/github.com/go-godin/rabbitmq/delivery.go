package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Delivery struct {
	amqp.Delivery
}

func (d Delivery) NackDelivery(multiple, requeue bool, topic string) error {
	var requeueVal string
	if requeue {
		requeueVal = "1"
	} else {
		requeueVal = "0"
	}

	nackCounter.With("routing_key", topic, "requeue", requeueVal).Add(1)
	return d.Nack(multiple, requeue)
}

func (d Delivery) IncrementTransportErrorCounter(topic string) {
	transportError.With("routing_key", topic).Add(1)
}

func (d Delivery) DecrementTransportErrorCounter() {
	transportError.With("routing_key", d.RoutingKey).Add(-1)
}
