package jwtutil

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Sign(claims jwt.Claims, key string) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
}

func Parse(claims jwt.Claims, token string, key string) (jwt.Claims, error) {
	var jwtToken *jwt.Token
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	if jwtToken != nil && jwtToken.Valid {
		return jwtToken.Claims, nil
	}
	return nil, fmt.Errorf("invalid jwt token")
}
