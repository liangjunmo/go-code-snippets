package tracing

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/stretchr/testify/require"
)

func TestGRPCUnaryServerInterceptor(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"
	key2 := "key2"
	value2 := "value2"

	SetTracingIDKey(key)
	SetTracingIDGenerator(func() string {
		return value
	})
	AppendTracingKeys([]string{key2})

	md := metadata.New(map[string]string{
		key2: value2,
	})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	handler := func(ctx context.Context, req any) (any, error) {
		require.Equal(t, value, ctx.Value(key))
		require.Equal(t, value2, ctx.Value(key2))
		return nil, nil
	}
	_, err := GRPCUnaryServerInterceptor(ctx, nil, nil, handler)
	require.Nil(t, err)
}

func TestGRPCUnaryClientInterceptor(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"
	key2 := "key2"
	value2 := "value2"

	SetTracingIDKey(key)
	SetTracingIDGenerator(func() string {
		return value
	})
	AppendTracingKeys([]string{key2})

	md := metadata.New(map[string]string{
		key:  "",
		key2: "",
	})
	ctx := context.WithValue(context.Background(), key, value)
	ctx = context.WithValue(ctx, key2, value2)
	ctx = metadata.NewOutgoingContext(ctx, md)

	handler := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		require.Equal(t, value, ctx.Value(key))
		require.Equal(t, value2, ctx.Value(key2))
		return nil
	}
	err := GRPCUnaryClientInterceptor(ctx, "", nil, nil, nil, handler, nil)
	require.Nil(t, err)
}
