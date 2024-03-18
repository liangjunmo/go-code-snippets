package rsautil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPrivateKeyWithFile(t *testing.T) {
	_, err := NewPrivateKeyWithFile("./testdata/private.pem")
	require.Nil(t, err)
}

func TestNewPublicKeyWithFile(t *testing.T) {
	_, err := NewPublicKeyWithFile("./testdata/public.pem")
	require.Nil(t, err)
}
