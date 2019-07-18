// Code generated by Godin v0.4.0; DO NOT EDIT.
package amqp

import (
	"github.com/pkg/errors"

	"github.com/go-godin/log"
	middleware "github.com/go-godin/middleware/amqp"
	"github.com/go-godin/rabbitmq"

	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/subscriber"
)

type SubscriberSet struct {
	userCreatedSubscriber rabbitmq.Subscriber
	userDeletedSubscriber rabbitmq.Subscriber
}

// Subscriptions will initialize all AMQP subscriptions. For each subscription, a new AMQP channel is created on which
// the Consume() goroutine will run (one channel per thread policy).
func Subscriptions(conn *rabbitmq.RabbitMQ) SubscriberSet {
	userCreatedSubscriberChannel, _ := conn.NewChannel()
	userDeletedSubscriberChannel, _ := conn.NewChannel()

	return SubscriberSet{
		userCreatedSubscriber: rabbitmq.NewSubscriber(userCreatedSubscriberChannel, &rabbitmq.Subscription{
			Exchange: "exchange-name",
			Topic:    "user.created",
			AutoAck:  false,
			Queue: rabbitmq.SubscriptionQueue{
				AutoDelete: false,
				Durable:    true,
				Exclusive:  false,
				Name:       "user-created-queue",
				NoWait:     false,
			},
		}),

		userDeletedSubscriber: rabbitmq.NewSubscriber(userDeletedSubscriberChannel, &rabbitmq.Subscription{
			Exchange: "exchange-name",
			Topic:    "user.deleted",
			AutoAck:  false,
			Queue: rabbitmq.SubscriptionQueue{
				AutoDelete: false,
				Durable:    true,
				Exclusive:  false,
				Name:       "user-deleted-queue",
				NoWait:     false,
			},
		}),
	}
}

// UserCreatedSubscriber sets up the subscription to the 'user.created' topic. All middleware is automatically registered
// and called in the following order: RequestID => PrometheusInstrumentation => Logging => subscriber.UserCreatedSubscriber
// The RequestID middlware will extract the requestId from the delivery header or generate a new one. The requestId is
// then mad available through the context.
func (ss SubscriberSet) UserCreatedSubscriber(logger log.Logger, usecase service.Stringer) error {
	handler := subscriber.UserCreatedSubscriber(logger, usecase)
	handler = middleware.Logging(logger, "user.created", handler)
	handler = middleware.PrometheusInstrumentation("user.created", handler)
	handler = middleware.RequestID(handler)

	if err := ss.userCreatedSubscriber.Subscribe(handler); err != nil {
		logger.Error("failed to subscribe to user.created", "err", err, "transport", "AMQP")
		return errors.Wrap(err, "failed to subscribe to user.created")
	}
	logger.Info(
		"subscribed to topic 'user.created'",
		"topic", ss.userCreatedSubscriber.Subscription.Topic,
		"queue", ss.userCreatedSubscriber.Subscription.Queue.Name,
		"exchange", ss.userCreatedSubscriber.Subscription.Exchange,
		"transport", "AMQP",
	)

	return nil
}

// UserDeletedSubscriber sets up the subscription to the 'user.deleted' topic. All middleware is automatically registered
// and called in the following order: RequestID => PrometheusInstrumentation => Logging => subscriber.UserDeletedSubscriber
// The RequestID middlware will extract the requestId from the delivery header or generate a new one. The requestId is
// then mad available through the context.
func (ss SubscriberSet) UserDeletedSubscriber(logger log.Logger, usecase service.Stringer) error {
	handler := subscriber.UserDeletedSubscriber(logger, usecase)
	handler = middleware.Logging(logger, "user.deleted", handler)
	handler = middleware.PrometheusInstrumentation("user.deleted", handler)
	handler = middleware.RequestID(handler)

	if err := ss.userDeletedSubscriber.Subscribe(handler); err != nil {
		logger.Error("failed to subscribe to user.deleted", "err", err, "transport", "AMQP")
		return errors.Wrap(err, "failed to subscribe to user.deleted")
	}
	logger.Info(
		"subscribed to topic 'user.deleted'",
		"topic", ss.userDeletedSubscriber.Subscription.Topic,
		"queue", ss.userDeletedSubscriber.Subscription.Queue.Name,
		"exchange", ss.userDeletedSubscriber.Subscription.Exchange,
		"transport", "AMQP",
	)

	return nil
}

// Shutdown will call Shutdown() on all registered subscriptions
func (ss SubscriberSet) Shutdown() (err error) {
	err = ss.userCreatedSubscriber.Shutdown()
	err = ss.userDeletedSubscriber.Shutdown()

	return err
}
