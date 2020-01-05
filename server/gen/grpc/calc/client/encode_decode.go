// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/mdetweil/demo-reverse-proxy/server/design

package client

import (
	"context"

	calc "github.com/mdetweil/demo-reverse-proxy/server/gen/calc"
	calcpb "github.com/mdetweil/demo-reverse-proxy/server/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildAddFunc builds the remote method to invoke for "calc" service "add"
// endpoint.
func BuildAddFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*calcpb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &calcpb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to calc add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calc.AddPayload", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the calc add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildMultiplyFunc builds the remote method to invoke for "calc" service
// "multiply" endpoint.
func BuildMultiplyFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Multiply(ctx, reqpb.(*calcpb.MultiplyRequest), opts...)
		}
		return grpccli.Multiply(ctx, &calcpb.MultiplyRequest{}, opts...)
	}
}

// EncodeMultiplyRequest encodes requests sent to calc multiply endpoint.
func EncodeMultiplyRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.MultiplyPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
	}
	return NewMultiplyRequest(payload), nil
}

// DecodeMultiplyResponse decodes responses from the calc multiply endpoint.
func DecodeMultiplyResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.MultiplyResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyResponse", v)
	}
	res := NewMultiplyResult(message)
	return res, nil
}
