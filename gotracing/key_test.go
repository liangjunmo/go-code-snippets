package gotracing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetTracingIDKey(t *testing.T) {
	resetTracingKeys()
	key := "key"
	SetTracingIDKey(key)
	require.Equal(t, key, tracingIDKey)
}

func TestAppendTracingKeys(t *testing.T) {
	resetTracingKeys()
	key := "key"
	AppendTracingKeys([]string{key})
	require.Equal(t, []string{defaultTracingIDKey, key}, tracingKeys)
}
