// Code generated by Godin vv0.5.0; DO NOT EDIT.

package endpoint

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/go-godin/log"
	"github.com/go-godin/middleware"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

type Set struct {
	HelloEndpoint endpoint.Endpoint
}

func Endpoints(service service.Stringer, logger log.Logger) Set {

	var hello endpoint.Endpoint
	{
		hello = HelloEndpoint(service)
		hello = middleware.InstrumentGRPC("Hello")(hello)
		hello = middleware.Logging(logger, "Hello")(hello)
		hello = middleware.RequestID()(hello)
	}

	return Set{
		HelloEndpoint: hello,
	}
}
