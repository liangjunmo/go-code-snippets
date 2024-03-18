package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrace(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"

	SetTracingIDKey(key)
	SetTracingIDGenerator(func() string {
		return value
	})

	ctx := Trace(context.Background())
	require.Equal(t, value, ctx.Value(key))
}
