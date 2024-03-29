// Code generated by Godin vv0.4.0; DO NOT EDIT.
package amqp

import (
	"github.com/go-godin/log"
	amqpMiddleware "github.com/go-godin/middleware/amqp"
	"github.com/go-godin/rabbitmq"
	"github.com/lukasjarosch/godin-examples/hello/internal/service"
	"github.com/lukasjarosch/godin-examples/hello/internal/service/subscriber"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type SubscriberSet struct {
	UserCreated rabbitmq.Subscriber
}

func Subscriptions(channel *amqp.Channel) SubscriberSet {
	return SubscriberSet{
		UserCreated: rabbitmq.NewSubscriber(channel, &rabbitmq.Subscription{
			AutoAck:  false,
			Exchange: "exchange-name",
			Topic:    "user.created",
			Queue: rabbitmq.SubscriptionQueue{
				Name:       "user-created-queue",
				NoWait:     false,
				AutoDelete: false,
				Durable:    true,
				Exclusive:  false,
			},
		}),
	}
}

func (ss SubscriberSet) UserCreatedSubscription(logger log.Logger, usecase service.Hello) error {
	userCreatedHandler := subscriber.UserCreatedHandler(logger, usecase)
	userCreatedHandler = amqpMiddleware.Logging(logger, "user.created", userCreatedHandler)
	userCreatedHandler = amqpMiddleware.PrometheusInstrumentation("user.created", userCreatedHandler)
	userCreatedHandler = amqpMiddleware.RequestID(userCreatedHandler)

	if err := ss.UserCreated.Subscribe(userCreatedHandler); err != nil {
		logger.Error("failed to subscribe to user.created", "err", err, "transport", "AMQP")
		return errors.Wrap(err, "failed to subscribe to user.created")
	}
	logger.Info(
		"subscribed to topic 'user.created'",
		"topic", ss.UserCreated.Subscription.Topic,
		"queue", ss.UserCreated.Subscription.Queue.Name,
		"exchange", ss.UserCreated.Subscription.Exchange,
		"transport", "AMQP",
	)

	return nil
}
