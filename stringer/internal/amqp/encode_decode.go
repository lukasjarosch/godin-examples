package amqp

import (
	pb "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/rs/xid"
)

// UserCreatedEncoder is called just before publishing an event to 'user.created' and encodes
// the DAO to protobuf.
func UserCreatedEncoder(event interface{}) (*pb.UserCreatedEvent, error) {
	var encoded pb.UserCreatedEvent

	encoded.Id = xid.New().String()
	encoded.Email = "some-mail@gmail.com"

	return &encoded, nil
}
