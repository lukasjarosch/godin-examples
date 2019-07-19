package rabbitmq

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const EnvironmentVariableName = "AMQP_ADDRESS"

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel // used for setup
	URI        string
}

func NewRabbitMQ(connectionString string) *RabbitMQ {
	return &RabbitMQ{
		URI: connectionString,
	}
}

func NewRabbitMQFromEnv() (*RabbitMQ, error) {
	connection := os.Getenv(EnvironmentVariableName)
	if connection == "" {
		return nil, fmt.Errorf("missing AMQP connection string, set %s env variable", EnvironmentVariableName)
	}
	return NewRabbitMQ(connection), nil
}

func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial(r.URI)
	if err != nil {
		return err
	}
	r.Connection = conn

	return nil
}

// NewChannel creates a new amqp.Channel on the current connection
func (r *RabbitMQ) NewChannel() (channel *amqp.Channel, err error) {
	if r.Connection == nil {
		return nil, fmt.Errorf("cannot create channel without a connection")
	}
	channel, err = r.Connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open AMQP channel")
	}

	return channel, nil
}

func (r *RabbitMQ) DeclareExchange(name, typ string, durable, autoDelete, internal, noWait bool) (err error) {
	if err := r.Channel.ExchangeDeclare(name, typ, durable, autoDelete, internal, noWait, nil); err != nil {
		return errors.Wrap(err, "failed to delcare RabbitMQ exchange")
	}
	return nil
}

func (r *RabbitMQ) Close() error {
	return r.Connection.Close()
}
