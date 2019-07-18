package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/domain"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/endpoint"
)

// ----------------[ ERRORS ]----------------

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
	switch err {
	case domain.ErrNotImplemented:
		return status.Error(codes.Unimplemented, err.Error())
	case domain.ErrUnauthenticated:
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
	return err
}

// ----------------[ MAPPING FUNCS ]----------------

// TODO: this is a nice spot for convenience mapping functions :)

// ----------------[ ENCODER / DECODER ]----------------
// HelloRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level HelloRequest
func HelloRequestDecoder(pbRequest *pb.HelloRequest) (request endpoint.HelloRequest, err error) {
	request = endpoint.HelloRequest{
		Name: "Hans Peter",
	}
	return request, err
}

// HelloResponseEncoder encodes the domain-level HelloResponse into a protobuf HelloResponse
func HelloResponseEncoder(response endpoint.HelloResponse) (pbResponse *pb.HelloResponse, err error) {
	pbResponse = &pb.HelloResponse{}
	return pbResponse, err
}

// HelloRequestEncoder encodes the domain-level HelloRequest into a protobuf HelloRequest
func HelloRequestEncoder(request endpoint.HelloRequest) (pbRequest *pb.HelloRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return
	return pbRequest, err
}

// HelloResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level HelloResponse
func HelloResponseDecoder(pbResponse *pb.HelloResponse) (response endpoint.HelloResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return
	return response, err
}
