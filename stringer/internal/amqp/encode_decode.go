package amqp

import (
	"fmt"

	"github.com/go-godin/rabbitmq"

	"github.com/golang/protobuf/proto"
	pb "github.com/lukasjarosch/godin-examples/greeter/api"
	userCreatedProto "github.com/lukasjarosch/godin-examples/greeter/api"
	userDeletedProto "github.com/lukasjarosch/godin-examples/greeter/api"
)

// UserCreatedEncoder is called just before publishing an event to 'user.created' and encodes
// the DAO to protobuf.
func UserCreatedEncoder(event interface{}) (*pb.UserCreatedEvent, error) {
	var encoded pb.UserCreatedEvent

	// TODO: map to protobuf

	return &encoded, nil
}

// UserCreatedDecoder will unmarshal the Body of the incoming Delivery into a protobuf message.
//
// Note: Godin will not regenerate this file, only append new functions. So if this file was already present when
// you added the subscriber, you need to fix the imports by adding:
//	 userCreatedProto "github.com/lukasjarosch/godin-examples/greeter/api"
func UserCreatedDecoder(delivery *rabbitmq.Delivery) (decoded interface{}, err error) {
	event := &userCreatedProto.UserCreatedEvent{}

	if err := proto.Unmarshal(delivery.Body, event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Delivery into *userCreatedProto.UserCreatedEvent: %s", err)
	}

	return event, nil
}

// UserDeletedDecoder will unmarshal the Body of the incoming Delivery into a protobuf message.
//
// Note: Godin will not regenerate this file, only append new functions. So if this file was already present when
// you added the subscriber, you need to fix the imports by adding:
//	 userDeletedProto "github.com/lukasjarosch/godin-examples/greeter/api"
func UserDeletedDecoder(delivery *rabbitmq.Delivery) (decoded interface{}, err error) {
	event := &userDeletedProto.UserCreatedEvent{}

	if err := proto.Unmarshal(delivery.Body, event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Delivery into *userDeletedProto.UserCreatedEvent: %s", err)
	}

	return event, nil
}
