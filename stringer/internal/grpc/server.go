// Code generated by Godin vv0.4.0; DO NOT EDIT.

package grpc

import (
	"context"

	kitGrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/go-godin/log"
	pb "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/endpoint"
)

type grpcServer struct {
	HelloHandler kitGrpc.Handler
}

func NewServer(endpoints endpoint.Set, logger log.Logger) pb.GreeterServiceServer {
	// TODO: configurable ServerOptions via godin.json
	options := []kitGrpc.ServerOption{}

	return &grpcServer{
		HelloHandler: kitGrpc.NewServer(
			endpoints.HelloEndpoint,
			DecodeHelloRequest,
			EncodeHelloResponse,
			append(options)...,
		),
	}
}

func (s *grpcServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, resp, err := s.HelloHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, EncodeError(err)
	}
	return resp.(*pb.HelloResponse), nil
}
