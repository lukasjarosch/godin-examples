package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lukasjarosch/godin-examples/user/api"
	"github.com/lukasjarosch/godin-examples/user/internal/service/endpoint"
	"github.com/lukasjarosch/godin-examples/user/internal/service"
)

// ----------------[ ERRORS ]----------------

// EncodeError encodes domain-level errors into gRPC transport-level errors
func EncodeError(err error) error {
	switch err {
	case service.ErrNotImplemented:
		return status.Error(codes.Unimplemented, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

// ----------------[ MAPPING FUNCS ]----------------

// TODO: this is a nice spot for convenience mapping functions :)

// ----------------[ ENCODER / DECODER ]----------------
// CreateRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level CreateRequest
func CreateRequestDecoder(pbRequest *pb.CreateRequest) (request endpoint.CreateRequest, err error) {
	request = endpoint.CreateRequest{
		Email: pbRequest.Email,
		Username: pbRequest.Name,
	}
	return request, nil
}

// CreateResponseEncoder encodes the domain-level CreateResponse into a protobuf CreateResponse
func CreateResponseEncoder(response endpoint.CreateResponse) (pbResponse *pb.CreateResponse, err error) {
	pbResponse = &pb.CreateResponse{
		User: &pb.User{
			Id: response.User.ID,
			Name: response.User.Name,
			Email: response.User.Email,
		},
	}
	return pbResponse, nil
}

// CreateRequestEncoder encodes the domain-level CreateRequest into a protobuf CreateRequest
func CreateRequestEncoder(request endpoint.CreateRequest) (pbRequest *pb.CreateRequest, err error) {
	pbRequest = &pb.CreateRequest{
		Name: request.Username,
		Email: request.Email,
	}
	return pbRequest, nil
}

// CreateResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level CreateResponse
func CreateResponseDecoder(pbResponse *pb.CreateResponse) (response endpoint.CreateResponse, err error) {
	response = endpoint.CreateResponse{
		User: &service.UserEntity{
			ID: pbResponse.User.Id,
			Email: pbResponse.User.Email,
			Name: pbResponse.User.Name,
		},
	}
	return response, nil
}

// GetRequestEncoder encodes the domain-level GetRequest into a protobuf GetRequest
func GetRequestEncoder(request endpoint.GetRequest) (pbRequest *pb.GetRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return
	return pbRequest, err
}


// GetResponseEncoder encodes the domain-level GetResponse into a protobuf GetResponse
func GetResponseEncoder(response endpoint.GetResponse) (pbResponse *pb.GetResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return
	return pbResponse, err
}


// GetRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level GetRequest
func GetRequestDecoder(pbRequest *pb.GetRequest) (request endpoint.GetRequest, err error) {
	// TODO: map 'pbRequest' to 'request' and return
	return request, err
}


// GetResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level GetResponse
func GetResponseDecoder(pbResponse *pb.GetResponse) (response endpoint.GetResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return
	return response, err
}

// ListRequestEncoder encodes the domain-level ListRequest into a protobuf ListRequest
func ListRequestEncoder(request endpoint.ListRequest) (pbRequest *pb.ListRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return
	return pbRequest, err
}


// ListResponseEncoder encodes the domain-level ListResponse into a protobuf ListResponse
func ListResponseEncoder(response endpoint.ListResponse) (pbResponse *pb.ListResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return
	return pbResponse, err
}


// ListRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level ListRequest
func ListRequestDecoder(pbRequest *pb.ListRequest) (request endpoint.ListRequest, err error) {
	// TODO: map 'pbRequest' to 'request' and return
	return request, err
}


// ListResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level ListResponse
func ListResponseDecoder(pbResponse *pb.ListResponse) (response endpoint.ListResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return
	return response, err
}

// DeleteRequestEncoder encodes the domain-level DeleteRequest into a protobuf DeleteRequest
func DeleteRequestEncoder(request endpoint.DeleteRequest) (pbRequest *pb.DeleteRequest, err error) {
	// TODO: map 'request' to 'pbRequest' and return
	return pbRequest, err
}


// DeleteResponseEncoder encodes the domain-level DeleteResponse into a protobuf DeleteResponse
func DeleteResponseEncoder(response endpoint.DeleteResponse) (pbResponse *pb.EmptyResponse, err error) {
	// TODO: map 'response' to 'pbResponse' and return
	return pbResponse, err
}


// DeleteRequestDecoder maps the protobuf request of the gRPC transport layer onto the domain-level DeleteRequest
func DeleteRequestDecoder(pbRequest *pb.DeleteRequest) (request endpoint.DeleteRequest, err error) {
	request = endpoint.DeleteRequest{
		Id:pbRequest.Id,
	}
	return request, nil
}


// DeleteResponseDecoder maps the protobuf response of the gRPC transport layer onto the domain-level DeleteResponse
func DeleteResponseDecoder(pbResponse *pb.EmptyResponse) (response endpoint.DeleteResponse, err error) {
	// TODO: map 'pbResponse' to 'response' and return
	return response, err
}
