package amqp

import (
	pb "github.com/lukasjarosch/godin-examples/greeter/api"
)

// UserCreatedEncoder is called just before publishing an event to 'user.created' and encodes
// the DAO to protobuf.
func UserCreatedEncoder(event interface{}) (*pb.UserCreatedEvent, error) {
	var encoded pb.UserCreatedEvent

	// TODO: map to protobuf

	return &encoded, nil
}
