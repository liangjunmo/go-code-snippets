package gotracing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetTracingIDKey(t *testing.T) {
	resetTracingKeys()
	SetTracingIDKey("test")
	require.Equal(t, "test", tracingIDKey)
}

func TestAppendTracingKeys(t *testing.T) {
	resetTracingKeys()
	AppendTracingKeys([]string{"test"})
	require.Equal(t, []string{defaultTracingIDKey, "test"}, tracingKeys)
}
