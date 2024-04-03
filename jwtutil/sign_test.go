package jwtutil

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

type Claims struct {
	jwt.StandardClaims
	UID int
}

func TestSignAndParse(t *testing.T) {
	uid := 1
	key := "secret"

	token, err := Sign(&Claims{UID: uid}, key)
	require.Nil(t, err)
	require.True(t, len(token) > 0)

	jwtClaims, err := Parse(&Claims{}, token, key)
	require.Nil(t, err)
	claims := jwtClaims.(*Claims)
	require.Equal(t, uid, claims.UID)
}
