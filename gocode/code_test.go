package gocode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCode(t *testing.T) {
	code := Code("code")
	require.Equal(t, "code", code.Error())
	require.Equal(t, "code", code.String())
}
