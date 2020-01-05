// Code generated by goa v3.0.9, DO NOT EDIT.
//
// world gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/mdetweil/demo-reverse-proxy/server/design

package client

import (
	"context"

	worldpb "github.com/mdetweil/demo-reverse-proxy/server/gen/grpc/world/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildHelloFunc builds the remote method to invoke for "world" service
// "hello" endpoint.
func BuildHelloFunc(grpccli worldpb.WorldClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Hello(ctx, reqpb.(*worldpb.HelloRequest), opts...)
		}
		return grpccli.Hello(ctx, &worldpb.HelloRequest{}, opts...)
	}
}

// DecodeHelloResponse decodes responses from the world hello endpoint.
func DecodeHelloResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*worldpb.HelloResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("world", "hello", "*worldpb.HelloResponse", v)
	}
	res := NewHelloResult(message)
	return res, nil
}