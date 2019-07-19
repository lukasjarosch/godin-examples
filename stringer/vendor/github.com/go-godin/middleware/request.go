package middleware

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"

	md2 "github.com/go-godin/grpc-metadata"
	"github.com/go-kit/kit/endpoint"
)

func RequestID() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			// gRPC metadata
			if md, ok := metadata.FromIncomingContext(ctx); ok {
				requestID := md.Get(string(md2.RequestID))
				if len(requestID) > 0 {
					return next(ctx, request)
				}

				md.Append(string(md2.RequestID), uuid.New().String())
				ctx = metadata.NewIncomingContext(ctx, md)
				return next(ctx, request)
			}

			// no metadata or context, at least ensure the requestId exists
			ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(string(md2.RequestID), uuid.New().String()))
			return next(ctx, request)
		}
	}
}
