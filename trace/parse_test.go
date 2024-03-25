package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	resetTraceKeys()

	key := "key"
	value := "value"

	SetTraceIDKey(key)

	ctx := context.WithValue(context.Background(), key, value)
	labels := Parse(ctx)
	require.Equal(t, value, labels[key])
}
