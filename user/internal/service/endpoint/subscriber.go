package endpoint

import (
	"github.com/go-godin/log"
	"github.com/go-godin/middleware"
	"github.com/go-kit/kit/endpoint"
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
)

type SubscriberSet struct {
	UserCreatedEndpoint endpoint.Endpoint
}

func Subscriptions(subscription usecase.Subscriber, logger log.Logger) SubscriberSet {
	userCreated := UserCreatedEndpoint(subscription)
	userCreated = middleware.InstrumentRabbitMQ("user.created")(userCreated)

	return SubscriberSet{
		UserCreatedEndpoint: userCreated,
	}
}
