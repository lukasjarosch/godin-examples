package amqp

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	kitAMQP "github.com/go-kit/kit/transport/amqp"
	"github.com/lukasjarosch/godin-examples/user/internal/service/endpoint"
	amqp2 "github.com/lukasjarosch/godin/pkg/amqp"
	"github.com/streadway/amqp"
)

type amqpSubscriptions struct {
	userCreated *kitAMQP.Subscriber
	channel     *amqp.Channel
}

// NopSubscriptionResponseEncoder is used since this AMQP implementation is not meant for RPC over AMQP but for
// fire-and-forget events, thus not needing to send a response.
var NopSubscriptionResponseEncoder = func(context.Context, *amqp.Publishing, interface{}) error { return nil }

func Subscriptions(endpoints endpoint.SubscriberSet, channel *amqp.Channel) *amqpSubscriptions {
	options := []kitAMQP.SubscriberOption{
		kitAMQP.SubscriberResponsePublisher(kitAMQP.NopResponsePublisher),
	}

	f := func(ctx context.Context, deliv *amqp.Delivery, ch kitAMQP.Channel, pub *amqp.Publishing) context.Context {
		logrus.Error(deliv.Type)
		return ctx
	}

	return &amqpSubscriptions{
		channel: channel,
		userCreated: kitAMQP.NewSubscriber(
			endpoints.UserCreatedEndpoint,
			DecodeUserCreatedEvent,
			NopSubscriptionResponseEncoder,
			append(
				options,
				kitAMQP.SubscriberErrorEncoder(ErrorEncoder),
				kitAMQP.SubscriberAfter(f),
			)...,
		),
	}
}

func (sub *amqpSubscriptions) SubscribeUserCreated() error {
	s := amqp2.NewSubscriber(sub.channel, &amqp2.Subscription{
		Queue: amqp2.SubscriptionQueue{
			Name:       "some-queue",
			Exclusive:  false,
			Durable:    true,
			AutoDelete: false,
			NoWait:     false,
		},
		Topic:    "user.created",
		Exchange: "exchange-name",
		AutoAck:  false,
	})
	err := s.Subscribe(func(delivery amqp.Delivery) {
		sub.userCreated.ServeDelivery(sub.channel)(&delivery)
		delivery.Ack(false)
	})
	if err != nil {
		return errors.Wrap(err, "failed to handle delivery")
	}
	return nil
}
