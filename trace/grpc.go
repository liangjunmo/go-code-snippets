package trace

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, exist := metadata.FromIncomingContext(ctx)
	for _, key := range traceKeys {
		if !exist {
			ctx = context.WithValue(ctx, key, "")
			continue
		}

		mdKey := strings.ToLower(key)
		if values, ok := md[mdKey]; !ok {
			ctx = context.WithValue(ctx, key, "")
		} else if len(values) == 0 {
			ctx = context.WithValue(ctx, key, "")
		} else {
			ctx = context.WithValue(ctx, key, values[0])
		}
	}
	if ctx.Value(traceIDKey) == "" {
		ctx = Trace(ctx)
	}
	return handler(ctx, req)
}

func GRPCUnaryClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, exist := metadata.FromOutgoingContext(ctx)
	if !exist {
		md = make(metadata.MD, len(traceKeys))
	}

	for _, key := range traceKeys {
		mdKey := strings.ToLower(key)
		if _, ok := md[mdKey]; ok {
			delete(md, mdKey)
		}

		value := ctx.Value(key)
		if value == nil {
			md[mdKey] = []string{""}
		} else {
			md[mdKey] = []string{value.(string)}
		}
	}

	ctx = metadata.NewOutgoingContext(ctx, md)
	return invoker(ctx, method, req, reply, cc, opts...)
}
