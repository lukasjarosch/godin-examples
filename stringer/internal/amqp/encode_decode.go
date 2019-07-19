package amqp

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/rs/xid"

	"github.com/go-godin/rabbitmq"

	pb "github.com/lukasjarosch/godin-examples/greeter/api"
)

// UserCreatedEncoder is called just before publishing an event to 'user.created' and encodes
// the DAO to protobuf.
func UserCreatedEncoder(event interface{}) (*pb.UserCreatedEvent, error) {
	var encoded pb.UserCreatedEvent

	encoded.Id = xid.New().String()
	encoded.Email = "hans@mail.com"

	return &encoded, nil
}

// UserCreatedDecoder
func UserCreatedDecoder(delivery *rabbitmq.Delivery) (decoded interface{}, err error) {
	event := &pb.UserCreatedEvent{}

	if err := proto.Unmarshal(delivery.Body, event); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal UserCreatedEvent")
	}

	return event, nil
}
