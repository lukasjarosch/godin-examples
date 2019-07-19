package rabbitmq

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type SubscriptionHandler func(ctx context.Context, delivery *Delivery)
type SubscriberDecoder func(delivery *Delivery) (decoded interface{}, err error)

// Subscription defines all data required to setup an AMQP Subscription
// All values, except the CTag are provided by the configuration or inferred by Godin.
type Subscription struct {
	Topic         string            `json:"topic" mapstructure:"topic"`
	Exchange      string            `json:"exchange" mapstructure:"exchange"`
	AutoAck       bool              `json:"auto_ack" mapstructure:"auto_ack"`
	CTag          string            `json:"-"` // generated
	Queue         SubscriptionQueue `json:"queue" mapstructure:"queue"`
	PrefetchCount int               `json:"prefetch_count" mapstructure:"prefetch_count"`
}

// SubscriptionQueue configures the queue on which the Subscription runs.
type SubscriptionQueue struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"no_wait"`
}

type handler struct {
	Implementation SubscriptionHandler
	done           chan error
}

// Subscriber handles AMQP subscriptions.
type Subscriber struct {
	channel      *amqp.Channel
	Subscription *Subscription
	Handler      handler
}

// NewSubscriber returns a new Subscriber with auto-generated CTag
func NewSubscriber(channel *amqp.Channel, subscription *Subscription) Subscriber {
	sub := strings.Replace(subscription.Topic, ".", "_", -1)
	ctag := fmt.Sprintf("%s_%s", sub, uuid.New().String())
	subscription.CTag = ctag

	return Subscriber{
		channel:      channel,
		Subscription: subscription,
	}
}

// Subscribe will declare the queue defined in the Subscription, bind it to the exchange and start consuming
// by calling the Handler in a goroutine.
func (c *Subscriber) Subscribe(handler SubscriptionHandler) error {
	err := c.channel.ExchangeDeclare(c.Subscription.Exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to declare exchange '%s'", c.Subscription.Exchange))
	}

	// ensure the queue exists
	queue, err := c.channel.QueueDeclare(
		c.Subscription.Queue.Name,
		c.Subscription.Queue.Durable,
		c.Subscription.Queue.AutoDelete,
		c.Subscription.Queue.Exclusive,
		c.Subscription.Queue.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	// establish exchange =topic=> queue binding
	if err = c.channel.QueueBind(
		queue.Name,
		c.Subscription.Topic,
		c.Subscription.Exchange,
		c.Subscription.Queue.NoWait,
		nil,
	); err != nil {
		return err
	}

	// channel prefetch-count
	if err := c.channel.Qos(c.Subscription.PrefetchCount, 0, false); err != nil {
		return err
	}

	// start consuming
	deliveries, err := c.channel.Consume(
		queue.Name,
		c.Subscription.CTag,
		c.Subscription.AutoAck,
		c.Subscription.Queue.Exclusive,
		false,
		c.Subscription.Queue.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	c.setHandler(handler)
	go c.handle(deliveries, c.Handler)

	return nil
}

// setHandler installs a SubscriptionHandler to use for this Subscription.
func (c *Subscriber) setHandler(handlerImpl SubscriptionHandler) {
	h := handler{
		Implementation: handlerImpl,
		done:           make(chan error),
	}
	c.Handler = h
}

// Handler is started by Subscribe() as Goroutine. For each received AMQP delivery,
// it will call the Implementation(delivery) to allow business logic for each delivery to run.
func (c *Subscriber) handle(deliveries <-chan amqp.Delivery, h handler) {
	for d := range deliveries {
		ctx := context.Background()
		h.Implementation(ctx, &Delivery{d})
	}
	h.done <- nil
}

// Shutdown will cancel the subscriber by it's CTag. It needs to be registered
// to a shutdown Handler.
func (c *Subscriber) Shutdown() error {
	if err := c.channel.Cancel(c.Subscription.CTag, true); err != nil {
		return err
	}
	<-c.Handler.done
	return nil
}
