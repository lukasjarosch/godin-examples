package amqp

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/lukasjarosch/godin-examples/user/internal/service/domain"
	"github.com/pkg/errors"

	"github.com/lukasjarosch/godin-examples/user/internal/service/endpoint"

	"github.com/golang/protobuf/proto"
	api "github.com/lukasjarosch/godin-examples/user/api"

	kitAMQP "github.com/go-kit/kit/transport/amqp"
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
	"github.com/streadway/amqp"
)

var ErrUnmarshalProtobuf = errors.New("failed to unmarshal message")
var ErrMarshalProtobuf = errors.New("failed to marshal message")

// ErrorEncoder is responsible of handling errors on subscriptions. Depending on the error
// a delivery can be ACKed, NACKed or REJECTed. By default, if no error occurred, the delivery
// should be ACKed. If an unknown error occured, the delivery is NACKed and requeued.
// This function should be registered as SubscriberErrorEncoder.
func ErrorEncoder(ctx context.Context, err error, deliv *amqp.Delivery, ch kitAMQP.Channel, pub *amqp.Publishing) {
	logrus.Error(err)

	switch err {
	case ErrUnmarshalProtobuf,
		domain.ErrMalformedEvent:
		deliv.Reject(false)
		break
	case domain.ErrSomethingFailed:
		deliv.Nack(false, true)
		break
	default:
		if deliv.Redelivered {
			deliv.Nack(false, false)
		}
		deliv.Nack(false, true)
	}

}

// EncodeUserCreatedEvent is used when publishing a 'user.created' event. It will marshall the given user
// into the UserCreated event protobuf message.
func EncodeUserCreatedEvent(ctx context.Context, publishing *amqp.Publishing, user interface{}) error {
	user2 := user.(*usecase.User)
	evt := &api.UserCreated{
		User: &api.User{
			Id:    user2.UserID,
			Email: user2.Email,
			Name:  user2.Username,
		},
	}

	buf, err := proto.Marshal(evt)
	if err != nil {
		return ErrMarshalProtobuf
	}

	publishing.Body = buf

	return nil
}

// DecodeUserCreatedEvent is used in the UserCreated subscriber to decode the incoming message into a usable DAO
func DecodeUserCreatedEvent(ctx context.Context, delivery *amqp.Delivery) (created interface{}, err error) {
	evt := &api.UserCreated{}
	if err := proto.Unmarshal(delivery.Body, evt); err != nil {
		return nil, ErrUnmarshalProtobuf
	}

	return nil, domain.ErrMalformedEvent
	if evt.User == nil {
		return nil, domain.ErrMalformedEvent
	}

	u := endpoint.UserCreatedEvent{
		User: &usecase.User{
			ID:       evt.User.Id,
			Email:    evt.User.Email,
			Username: evt.User.Name,
		},
	}

	return u, nil
}
