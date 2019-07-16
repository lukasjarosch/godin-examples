package subscriber

import (
	"github.com/go-godin/log"
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
	godinAMQP "github.com/lukasjarosch/godin/pkg/amqp"
	"github.com/streadway/amqp"
)

// UserCreatedSubscriber does ...
// Do NOT rename the function unless you know how Godin works.
func UserCreatedSubscriber(logger log.Logger, svc usecase.Service) godinAMQP.SubscriptionHandler {
	return func(delivery amqp.Delivery) {
		// TODO: handle user.created message
	}
}
