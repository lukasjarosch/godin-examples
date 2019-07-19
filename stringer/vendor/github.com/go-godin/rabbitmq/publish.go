package rabbitmq

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-godin/grpc-metadata"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

type Publishing struct {
	Topic           string `json:"topic" mapstructure:"topic"`
	Exchange        string `json:"exchange" mapstructure:"exchange"`
	DeliveryMode    uint8  `json:"delivery_mode" mapstructure:"delivery_mode"`
	ProtobufMessage string `json:"protobuf_message" mapstructure:"protobuf_message"` // unused in this package, workaround for Godin
}

type Publisher interface {
	Publish(ctx context.Context, event interface{}) error
}

type publisher struct {
	channel    *amqp.Channel
	Publishing *Publishing
}

// NewPublisher returns an AMQP publisher
func NewPublisher(channel *amqp.Channel, publishing *Publishing) publisher {
	return publisher{
		channel:    channel,
		Publishing: publishing,
	}
}

// Publish actually publishes the event
func (p publisher) Publish(ctx context.Context, event interface{}) error {
	err := p.channel.ExchangeDeclare(p.Publishing.Exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to declare exchange '%s'", p.Publishing.Exchange))
	}
	protobuf := event.(proto.Message)
	bodyBytes, err := proto.Marshal(protobuf)
	if err != nil {
		return err
	}

	// ensure the requestId is passed along
	requestId := grpc_metadata.GetRequestID(ctx)
	headers := amqp.Table{
		"requestId": requestId,
	}

	publishing := amqp.Publishing{
		Headers:      headers,
		ContentType:  "application/octet-stream",
		DeliveryMode: p.Publishing.DeliveryMode,
		Priority:     0,
		Body:         bodyBytes,
	}

	if err := p.channel.Publish(
		p.Publishing.Exchange,
		p.Publishing.Topic,
		false,
		false,
		publishing,
	); err != nil {
		return errors.Wrap(err, "failed to publish event")
	}
	return nil
}
