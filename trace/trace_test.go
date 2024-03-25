package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrace(t *testing.T) {
	resetTraceKeys()

	key := "key"
	value := "value"

	SetTraceIDKey(key)
	SetTraceIDGenerator(func() string {
		return value
	})

	ctx := Trace(context.Background())
	require.Equal(t, value, ctx.Value(key))
}
