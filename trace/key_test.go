package trace

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetTraceIDKey(t *testing.T) {
	resetTraceKeys()
	key := "key"
	SetTraceIDKey(key)
	require.Equal(t, key, traceIDKey)
}

func TestAppendTraceKeys(t *testing.T) {
	resetTraceKeys()
	key := "key"
	AppendTraceKeys([]string{key})
	require.Equal(t, []string{defaultTraceIDKey, key}, traceKeys)
}
