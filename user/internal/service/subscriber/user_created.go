package subscriber

import (
	"github.com/lukasjarosch/godin-examples/user/internal/service"
	godinAMQP "github.com/lukasjarosch/godin/pkg/amqp"
	"github.com/lukasjarosch/godin/pkg/log"
	"github.com/streadway/amqp"
)

// UserCreatedSubscriber does ...
// Do NOT rename the function unless you know how Godin works.
func UserCreatedSubscriber(logger log.Logger, svc service.User) godinAMQP.SubscriptionHandler {
	return func(delivery amqp.Delivery) {
		// TODO: handle user.created message
	}
}
