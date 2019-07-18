package endpoint

import (
    "context"
    "github.com/go-kit/kit/endpoint"

    "github.com/lukasjarosch/godin-examples/greeter/internal/service"
)

// HelloEndpoint provides service.Hello() as general endpoint
// which can be used with arbitrary transport layers.
func HelloEndpoint(service service.Greeter) endpoint.Endpoint {
    return func (ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(HelloRequest)
        _ = req // bypass "declared and not used" compiler error if the request is empty and not used
        greeting,err := service.Hello(ctx,req.Name,)

        return HelloResponse{
            Greeting: greeting,
            Err: err,
        }, err
    }
}