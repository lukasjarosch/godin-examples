package subscriber

import (
	"github.com/go-godin/log"
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
	godinAMQP "github.com/lukasjarosch/godin/pkg/amqp"
	"github.com/streadway/amqp"
)

func InitUserCreatedSubscriber(channel *amqp.Channel, svc usecase.Service, logger log.Logger) (subscriber godinAMQP.Subscriber, err error) {
	subscription := &godinAMQP.Subscription{
		Exchange: "exchange-name",
		AutoAck:  false,
		Queue: godinAMQP.SubscriptionQueue{
			AutoDelete: false,
			Durable:    true,
			Exclusive:  false,
			Name:       "some-queue",
			NoWait:     false,
		},
		Topic: "some.topic",
	}
	subscriber = godinAMQP.NewSubscriber(channel, subscription)
	if err = subscriber.Subscribe(UserCreatedSubscriber(logger, svc)); err != nil {
		logger.Error("failed to subscribe to user.created", "err", err)
		return subscriber, err
	}
	logger.Info("subscribed to topic 'user.created'", "topic", subscription.Topic, "queue", subscription.Queue.Name, "exchange", subscription.Exchange)
	return subscriber, nil
}
