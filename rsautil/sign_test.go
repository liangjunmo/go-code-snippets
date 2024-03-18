package rsautil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignAndVerify(t *testing.T) {
	privateKey, err := NewPrivateKeyWithFile("./testdata/private.pem")
	require.Nil(t, err)

	publicKey, err := NewPublicKeyWithFile("./testdata/public.pem")
	require.Nil(t, err)

	signature, err := SignWithSHA256(privateKey, "message")
	require.Nil(t, err)

	err = VerifySignatureWithSHA256(publicKey, signature, "message")
	require.Nil(t, err)
}
