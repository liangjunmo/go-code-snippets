package gotracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"

	SetTracingIDKey(key)

	ctx := context.WithValue(context.Background(), key, value)
	labels := Parse(ctx)
	require.Equal(t, value, labels[key])
}
